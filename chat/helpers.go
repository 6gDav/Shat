package chat

import (
	"encoding/json"
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
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

func BroadcastUserNameList() {
	var users = make(map[string]string)

	ClientsMu.RLock()
	for _, val := range Clients {
		if val.Conn != nil {
			users[val.IP] = val.Name
		}
	}
	ClientsMu.RUnlock()

	payload := WSResponse{
		Type: "USER_NAME_LIST",
		Data: users,
	}

	namesData, err := json.Marshal(payload)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("JSON marshalling error: " + err.Error()))
		})
		return
	}

	ClientsMu.RLock()
	defer ClientsMu.RUnlock()

	for ip, client := range Clients {
		if client.Conn != nil {
			err := client.Conn.WriteMessage(websocket.TextMessage, namesData)
			if err != nil {
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel("Failed to send list to IP " + ip + ": " + err.Error()))
				})
			}
		}
	}
}
