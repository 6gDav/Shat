package logic

import (
	"encoding/json"
	"fmt"

	"hosting_login_page/chat/client"
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func BroadcastUserNameList() {
	var users = make(map[string]string)

	client.ClientsMu.RLock()
	for _, val := range client.Clients {
		if val.Conn != nil {
			users[val.IP] = val.Name
		}
	}
	client.ClientsMu.RUnlock()

	payload := map[string]any{
		"type": "USER_NAME_LIST",
		"data": users,
	}

	namesData, err := json.Marshal(payload)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occurred while trying to convert the user list to a proper form " + err.Error()))
		})
		return
	}

	client.ClientsMu.RLock()
	defer client.ClientsMu.RUnlock()

	for ip, client := range client.Clients {
		if client.Conn != nil {
			err := client.Conn.WriteMessage(websocket.TextMessage, namesData)
			if err != nil {
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel(fmt.Sprintf("Failed to send user list to IP %s: %v", ip, err)))
				})
			}
		}
	}
}
