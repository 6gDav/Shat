package server

import (
	"encoding/json"
	"fmt"
	"hosting_login_page/chat"
	"hosting_login_page/logs"
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

	//webpage
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		chat.ClientsMu.Lock()
		if _, exists := chat.Clients[ip]; exists {
			http.Redirect(w, r, "/pages/Dasboard", http.StatusSeeOther)
		}
		chat.ClientsMu.Unlock()

		fileServer := http.FileServer(http.Dir("./hosted/build"))
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
		var data struct {
			Name string `json:"name"`
		}

		errdecode := json.NewDecoder(r.Body).Decode(&data)
		if errdecode != nil {
			logs.Logs.Add(widget.NewLabel("Failed to read JSON from IP address " + ip))
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		chat.ClientsMu.Lock()
		if _, exists := chat.Clients[ip]; !exists {
			chat.Clients[ip] = &chat.Client{
				IP:   ip,
				Name: data.Name,
				Conn: nil,
			}
			//Loging out
			logMsg := fmt.Sprintf("%+v\n", chat.Clients)
			logs.Logs.Add(widget.NewLabel("New user connected to teh server:  " + logMsg))
		}
		chat.ClientsMu.Unlock()

		//Return 200 (Ok)
		w.WriteHeader(http.StatusOK)

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

	portStart(mux)
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
