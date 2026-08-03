package chat

import (
	"hosting_login_page/logs"

	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type WSResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occurred while trying to set the Upgrader: " + err.Error()))
		})
		return
	}

	//managger
	ClientsMu.Lock()
	client, exists := Clients[ip]
	ClientsMu.Unlock()

	if !validateClientExistence(exists, conn) {
		return
	}

	//start and shotdown connection
	manageConnection(client, conn, ip)
}
