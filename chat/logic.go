package chat

import (
	"hosting_login_page/logs"

	"net"
	"net/http"

	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	ClientsMu.Lock()
	client, exists := Clients[ip]
	ClientsMu.Unlock()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Logs.Add(widget.NewLabel("Erro occured while trying to set the Upgrader: " + err.Error()))
		return
	}

	if !exists {
		errPayload := map[string]string{
			"type":    "error",
			"message": "Please submit your name first",
		}
		_ = conn.WriteJSON(errPayload)

		conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(4001, "Unauthorized"),
		)
		conn.Close()
		return
	}

	ClientsMu.Lock()
	client.Conn = conn
	ClientsMu.Unlock()

	logs.Logs.Add(widget.NewLabel("Client connected on this IP adderess: " + ip))

	defer func() {
		conn.Close()

		ClientsMu.Lock()
		if client, exists := Clients[ip]; exists {
			client.Conn = nil
		}
		ClientsMu.Unlock()

		logs.Logs.Add(widget.NewLabel("The conection was interrupted on this IP adderess: " + ip))
	}()

	for {
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			logs.Logs.Add(widget.NewLabel("Client diesconected on this IP adderess: " + ip))
			break
		}
	}
}
