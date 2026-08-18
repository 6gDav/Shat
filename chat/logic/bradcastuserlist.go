package logic

import (
	"encoding/json"
	"hosting_login_page/chat"
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func broadcastUserNameList() {
	var users = make(map[string]string)

	chat.ClientsMu.RLock()
	for _, val := range chat.Clients {
		if val.Conn != nil {
			users[val.IP] = val.Name
		}
	}
	chat.ClientsMu.RUnlock()

	payload := map[string]any{
		"type": "USER_NAME_LIST",
		"data": users,
	}

	namesData, err := json.Marshal(payload)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occured while trying to convert the user list to a proper form " + err.Error()))
		})
		return
	}

	chat.ClientsMu.RLock()
	defer chat.ClientsMu.RUnlock()

	for ip, client := range chat.Clients {
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
