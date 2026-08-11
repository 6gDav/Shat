package chat

import (
	"github.com/gorilla/websocket"
)

func validateClientExistence(exists bool, conn *websocket.Conn) bool {
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

		return false
	}
	return true
}

func generateRoomID(ip1, ip2 string) string {
	if ip1 < ip2 {
		return ip1 + "_" + ip2
	}
	return ip2 + "_" + ip1
}
