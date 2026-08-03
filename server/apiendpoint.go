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
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	mux.HandleFunc("GET /redirectuser", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := getIpAdress(r)
		if err != nil {
			http.Error(w, "Invalid IP address", http.StatusBadRequest)
			return
		}

		type RedirectResponse struct {
			Redirect bool   `json:"redirect"`
			UserName string `json:"userName,omitempty"`
		}

		var response RedirectResponse

		chat.ClientsMu.Lock()
		client, exists := chat.Clients[ip]
		if exists {
			response = RedirectResponse{
				Redirect: true,
				UserName: client.Name,
			}
		} else {
			response = RedirectResponse{
				Redirect: false,
			}
		}
		chat.ClientsMu.Unlock()

		if response.Redirect {
			logs.Logs.Add(widget.NewLabel("User redirected to dashboard IP: " + ip))
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
			return
		}
	})

	//submit user name
	mux.HandleFunc("POST /submit", func(w http.ResponseWriter, r *http.Request) {
		//Ip adress fetch
		ip, _, _ := getIpAdress(r)

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

	mux.HandleFunc("PATCH /submitnewusername", func(w http.ResponseWriter, r *http.Request) {
		//Ip adress fetch
		ip, _, err := getIpAdress(r)

		if err != nil {
			logMsg := fmt.Sprintf("Error occurred while this user %s tried to change username: %v", ip, err)
			logs.Logs.Add(widget.NewLabel(logMsg))
		}

		//Name fetch
		var data struct {
			Name string `json:"name"`
		}

		defer r.Body.Close()

		errdecode := json.NewDecoder(r.Body).Decode(&data)
		if errdecode != nil {
			logs.Logs.Add(widget.NewLabel("Failed to read JSON from IP address " + ip))
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		chat.ClientsMu.Lock()
		if client, exists := chat.Clients[ip]; exists {
			client.Name = data.Name
			logMessage := fmt.Sprintf("User name change on IP %s: The new name is: %s", ip, data.Name)
			logs.Logs.Add(widget.NewLabel(logMessage))
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

func getIpAdress(r *http.Request) (string, string, error) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	return ip, port, err
}
