package chat

import (
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func manageConnection(client *Client, conn *websocket.Conn, ip string) {
	ClientsMu.Lock()
	client.Conn = conn
	ClientsMu.Unlock()

	fyne.Do(func() {
		logs.Logs.Add(widget.NewLabel("Client connected on this IP address: " + ip))
	})

	BroadcastUserNameList()
	defer func() {
		conn.Close()

		ClientsMu.Lock()
		if c, ok := Clients[ip]; ok && c.Conn == conn {
			c.Conn = nil
		}
		ClientsMu.Unlock()

		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("The connection was interrupted on this IP address: " + ip))
			BroadcastUserNameList()
		})
	}()
	manageChat(conn, ip)
}

func manageChat(conn *websocket.Conn, ip string) {
	defer conn.Close()

	for {
		//in
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Client disconnected on this IP address: " + ip))
			})
			break
		}

		targetIP := msg["target_ip"]
		text := msg["text"]

		//out
		outMsg := map[string]string{
			"from":    ip,
			"to":      targetIP,
			"text":    text,
			"room_id": generateRoomID(ip, targetIP),
		}

		ClientsMu.RLock()
		senderClient := Clients[ip]
		targetClient := Clients[targetIP]
		ClientsMu.RUnlock()

		if targetClient != nil && targetClient.Conn != nil {
			err := targetClient.Conn.WriteJSON(outMsg)
			if err != nil {
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel("Error occurred while trying to send message: " + err.Error()))
				})
			}
		}

		if senderClient != nil && senderClient.Conn != nil {
			_ = senderClient.Conn.WriteJSON(outMsg)
		}
	}
}

func generateRoomID(ip1, ip2 string) string {
	if ip1 < ip2 {
		return ip1 + "_" + ip2
	}
	return ip2 + "_" + ip1
}

// func BroadcastMessage(msg map[string]string) {
// 	ClientsMu.Lock()
// 	defer ClientsMu.Unlock()

// 	for _, client := range Clients {
// 		if client.Conn != nil {
// 			_ = client.Conn.WriteJSON(msg)
// 		}
// 	}
// }
