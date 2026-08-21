package components

import (
	"hosting_login_page/chat/client"
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
		logs.Logs.Add(widget.NewLabel("DNS server successfully shut down..."))
	}

	client.ClientsMu.Lock()
	for _, client := range client.Clients {
		if client.Conn != nil {

			client.Conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseGoingAway, "Server shutting down"),
				time.Now().Add(time.Second),
			)
			client.Conn.Close()
			client.Conn = nil
		}
	}
	client.ClientsMu.Unlock()
}
