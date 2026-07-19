package components

import (
	"context"
	"fmt"
	"hosting_login_page/server"
	"log"
	"time"

	"fyne.io/fyne/v2/widget"
)

var clicked bool = false

func StartButtonElement() *widget.Button {
	//initialText := "Start"

	var button *widget.Button

	button = widget.NewButton("Start", func() {
		if clicked {
			clicked = false
			button.SetText("Start")
			stopServer()
		} else {
			clicked = true
			button.SetText("Stop")
			startServer()
		}
	})

	return button
}

func startServer() {
	server.SetAPIendpoint()
	server.SetMDNSserver()
}

func stopServer() {
	if server.MdnsServer != nil {
		server.MdnsServer.Shutdown()
		server.MdnsServer = nil
		fmt.Println("mDNS server sucessfully stopped")
	}

	if server.HttpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.HttpServer.Shutdown(ctx); err != nil {
			log.Printf("Error occured while trying to stop the server: %v", err)
		} else {
			fmt.Println("HTTP server sucessfully stopped")
		}
		server.HttpServer = nil
	}
}
