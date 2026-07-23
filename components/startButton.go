package components

import (
	"context"
	"fmt"
	"hosting_login_page/chat"
	"hosting_login_page/server"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

var Port int

func StartButtonElement() *widget.Button {

	val, err := strconv.Atoi(portInput.Text)
	if err != nil {
		log.Fatalf("Post input is not valid")
	}
	Port = val

	var clicked bool = false
	var button *widget.Button
	//initialText := "Start"

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
	//Closing mDNS service
	if server.MdnsServer != nil {
		server.MdnsServer.Shutdown()
		server.MdnsServer = nil
		fmt.Println("mDNS server sucessfully stopped")
	}

	//Closing WebSocket connections

	// chat.ClientsMu.Lock()
	// for _, client := range chat.Clients {
	// 	if client.Conn != nil {
	// 		client.Conn.Close()
	// 		client.Conn = nil
	// 	}
	// 	fmt.Println("All WebSocket connections closed")
	// }
	// chat.ClientsMu.Unlock()

	chat.ClientsMu.Lock()
	for ip, client := range chat.Clients {
		if client.Conn != nil {

			client.Conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseGoingAway, "Server shutting down"),
				time.Now().Add(time.Second),
			)
			client.Conn.Close()
			client.Conn = nil
			fmt.Println("WS close:", ip)
		}
	}
	chat.ClientsMu.Unlock()

	//Closing HTTP connections
	if server.HttpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := server.HttpServer.Shutdown(ctx); err != nil {
			log.Printf("Error occured while trying to stop the server: %v", err)
		} else {
			fmt.Println("HTTP server sucessfully stopped")
		}
		server.HttpServer = nil
	}
}
