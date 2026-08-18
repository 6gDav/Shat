package logic

import (
	"hosting_login_page/chat"
	"hosting_login_page/chat/helper"
	"hosting_login_page/logs"

	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	// Get user IP address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occurred while trying to set the Upgrader: " + err.Error()))
		})
		return
	}

	chat.ClientsMu.Lock()
	client, exists := chat.Clients[ip]

	if !exists {
		chat.ClientsMu.Unlock()
		helper.ValidateClientExistence(false, conn) // Reject if client is not registered
		return
	}

	// Close previous connection if it exists
	if client.Conn != nil {
		_ = client.Conn.Close()
	}

	// Update connection in the struct AND write it back to the map!
	client.Conn = conn
	chat.Clients[ip] = client
	chat.ClientsMu.Unlock()

	// Initialize and start the chat session
	chatManager := NewClientSession(client, conn, ip)
	//chatManager.SendChatHistory()
	chatManager.ManageConnection()
}
