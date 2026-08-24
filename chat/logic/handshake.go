package logic

import (
	"Shat/chat/client"
	"Shat/chat/helper"
	"Shat/logs"

	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Unable to start live connection: " + err.Error()))
		})
		return
	}

	client.ClientsMu.Lock()
	user, exists := client.Clients[ip]

	if !exists {
		client.ClientsMu.Unlock()
		helper.ValidateClientExistence(false, conn)
		return
	}

	if user.Conn != nil {
		_ = user.Conn.Close()
	}

	user.Conn = conn
	client.Clients[ip] = user
	client.ClientsMu.Unlock()

	chatManager := NewClientSession(user, conn, ip)
	chatManager.ManageConnection()
}
