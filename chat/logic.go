package chat

import (
	"fmt"

	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	ClientsMu.Lock()
	client, exists := Clients[ip]
	ClientsMu.Unlock()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
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

	fmt.Println("WebSocket connected for IP:", ip)

	defer func() {
		conn.Close()

		ClientsMu.Lock()
		if client, exists := Clients[ip]; exists {
			client.Conn = nil
		}
		ClientsMu.Unlock()

		fmt.Println("The conection was interrupted:", ip)
	}()

	for {
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Cliend disconected")
			break
		}
	}
}
