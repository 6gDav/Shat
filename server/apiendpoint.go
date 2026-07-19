package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"hosting_login_page/chat"
)

var HttpServer *http.Server

func SetAPIendpoint() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	mux.HandleFunc("POST /submit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Post hapend")
		//Ip adress fetch
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {

			ip = r.RemoteAddr
		}
		fmt.Println("Ip adress" + ip)
		//Name fetch
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the name")
		}
		defer r.Body.Close()

		if _, letezik := chat.Users[ip]; !letezik {
			chat.Users[ip] = string(body)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println("Your ip adress " + ip)
		fmt.Println("Your name " + string(body))

		fmt.Printf("%+v\n", chat.Users)
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
