package server

import (
	"fmt"
	"log"
	"net/http"
)

var HttpServer *http.Server

func SetAPIendpoint() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fileServer.ServeHTTP(w, r)
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
