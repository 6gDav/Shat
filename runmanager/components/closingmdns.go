package components

import (
	"hosting_login_page/chat"
	"hosting_login_page/logs"
	"hosting_login_page/server"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func ClosingmDNS() {
	if server.MdnsServer != nil {
		server.MdnsServer.Shutdown()
		server.MdnsServer = nil
		logs.Logs.Add(widget.NewLabel("mDNS Server successfully shot down..."))
	}

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
}
