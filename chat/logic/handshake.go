package logic

import (
	"hosting_login_page/chat/client"
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

	client.ClientsMu.Lock()
	user, exists := client.Clients[ip]

	if !exists {
		client.ClientsMu.Unlock()
		helper.ValidateClientExistence(false, conn) // Reject if client is not registered
		return
	}

	// Close previous connection if it exists
	if user.Conn != nil {
		_ = user.Conn.Close()
	}

	// Update connection in the struct AND write it back to the map!
	user.Conn = conn
	client.Clients[ip] = user
	client.ClientsMu.Unlock()

	// Initialize and start the chat session
	chatManager := NewClientSession(user, conn, ip)
	//chatManager.SendChatHistory()
	chatManager.ManageConnection()
}
