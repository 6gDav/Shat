package components

import (
	"context"
	"fmt"
	"hosting_login_page/chat"
	"hosting_login_page/logs"
	"hosting_login_page/server"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func StartButtonElement() *widget.Button {
	var clicked bool = false
	var button *widget.Button
	//initialText := "Start"

	button = widget.NewButton("Start", func() {
		if clicked {
			logs.Logs.Add(widget.NewLabel("Servers stopped..."))
			clicked = false
			button.SetText("Start")
			stopServer()
		} else {
			logs.Logs.Add(widget.NewLabel("Servers started..."))
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
		logs.Logs.Add(widget.NewLabel("mDNS Server successfully shot down..."))
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
			logs.Logs.Add(widget.NewLabel("Closing WebSocket connection on this ip adress: " + ip))
		}
	}
	chat.ClientsMu.Unlock()

	//Closing HTTP connections
	if server.HttpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := server.HttpServer.Shutdown(ctx); err != nil {
			errMsg := fmt.Sprintf("Error occured while trying to stop the server: %v", err)
			logs.Logs.Add(widget.NewLabel(errMsg))
		} else {
			logs.Logs.Add(widget.NewLabel("HTTP serves succesfully shot down..  "))
		}
		server.HttpServer = nil
	}
}
