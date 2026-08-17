package chat

import (
	"encoding/json"
	"fmt"
	"hosting_login_page/logs"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type Client struct {
	IP   string          `json:"ip"`
	Name string          `json:"name"`
	Conn *websocket.Conn `json:"-"`
}

func (c *Client) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{IP: %s, Name: %s}", c.IP, c.Name)
}

var (
	Clients   = make(map[string]*Client)
	ClientsMu sync.RWMutex
)

func BroadcastUserNameList() {
	var users = make(map[string]string)

	ClientsMu.RLock()
	for _, val := range Clients {
		if val.Conn != nil {
			users[val.IP] = val.Name
		}
	}
	ClientsMu.RUnlock()

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
