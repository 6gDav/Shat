package server

import (
	"fmt"
	"hosting_login_page/chat"
	"io"
	"log"
	"net"
	"net/http"

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
			log.Printf("Failed to read the name")
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
		fmt.Println("Your ip adress " + ip)
		fmt.Println("Your name " + string(body))

		chat.ClientsMu.RLock()
		fmt.Println("Heres the HashMap")
		fmt.Printf("%+v\n", chat.Clients)
		chat.ClientsMu.RUnlock()
	})

	//HandShaking
	//! Refactor this to the chat folder and solve the error problem (HTTP required WS)
	mux.HandleFunc("GET /setws", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		chat.ClientsMu.Lock()
		client, exists := chat.Clients[ip]
		chat.ClientsMu.Unlock()

		if !exists {
			http.Error(w, "Please submit your name first", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}

		chat.ClientsMu.Lock()
		client.Conn = conn
		chat.ClientsMu.Unlock()

		fmt.Println("WebSocket connected for IP:", ip)

		defer func() {
			conn.Close()

			chat.ClientsMu.Lock()
			if client, exists := chat.Clients[ip]; exists {
				client.Conn = nil
			}
			chat.ClientsMu.Unlock()

			fmt.Println("The conection was interrupted:", ip)
		}()

		for {
			var msg map[string]string
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("Cliend disconected")
				break
			}
		}
	})

	HttpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		fmt.Printf("Server is running on port %d\n", port)
		if err := HttpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Error occurred while trying to start the server: %v", err)
		}
	}()
}
