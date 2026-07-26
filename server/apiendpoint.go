package server

import (
	"encoding/json"
	"fmt"
	"hosting_login_page/chat"
	"hosting_login_page/logs"
	"io"
	"net"
	"net/http"

	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var HttpServer *http.Server

func SetAPIendpoint() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	//webpage
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	//submit user name
	mux.HandleFunc("POST /submit", func(w http.ResponseWriter, r *http.Request) {
		//Ip adress fetch
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {

			ip = r.RemoteAddr
		}

		//Name fetch
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logs.Logs.Add(widget.NewLabel("Faild to read the name on this IP address " + ip))
		}
		defer r.Body.Close()

		if _, exists := chat.Clients[ip]; !exists {
			name := string(body)

			chat.ClientsMu.Lock()
			chat.Clients[ip] = &chat.Client{
				IP:   ip,
				Name: name,
				Conn: nil,
			}
			chat.ClientsMu.Unlock()
		}

		w.WriteHeader(http.StatusOK)

		chat.ClientsMu.RLock()
		logMsg := fmt.Sprintf("%+v\n", chat.Clients)
		logs.Logs.Add(widget.NewLabel("New user connected to teh server:  " + logMsg))
		chat.ClientsMu.RUnlock()
	})

	//HandShaking
	mux.HandleFunc("GET /setws", func(w http.ResponseWriter, r *http.Request) {
		chat.HandShake(w, r, &upgrader)
	})

	mux.HandleFunc("GET /getusers", func(w http.ResponseWriter, r *http.Request) {

		chat.ClientsMu.Lock()
		err := json.NewEncoder(w).Encode(chat.Clients)
		chat.ClientsMu.Unlock()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logs.Logs.Add(widget.NewLabel("Error occured while trying to send the user data: " + err.Error()))
		}
	})

	PortStart(mux)
}

// HttpServer = &http.Server{
// 	Addr:    fmt.Sprintf(":%d", Port),
// 	Handler: mux,
// }

// go func() {
// 	fmt.Printf("Server is running on port %d\n", Port)
// 	if err := HttpServer.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Printf("Error occurred while trying to start the server: %v", err)
// 	}
// }()
