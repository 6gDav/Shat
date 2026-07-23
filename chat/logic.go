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

	if !exists {
		http.Error(w, "Please submit your name first", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
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
