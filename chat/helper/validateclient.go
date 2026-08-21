package helper

import "github.com/gorilla/websocket"

func ValidateClientExistence(exists bool, conn *websocket.Conn) bool {
	if !exists {
		errPayload := map[string]string{
			"type":    "error",
			"message": "Please submit your name first!",
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
