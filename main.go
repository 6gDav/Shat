package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./hosted/build")))

	fmt.Println("Server is runing on http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error occured while trying to start the server")
	}
}
