package server

import (
	"fmt"
	"log"
	"net/http"
)

func SetAPIendpoint() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	go func() {
		fmt.Printf("Server is running on port %d\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
		if err != nil {
			log.Fatalf("Error occurred while trying to start the server: %v", err)
		}
	}()
}
