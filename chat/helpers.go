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

func manageConnection(client *Client, conn *websocket.Conn, ip string) {
	ClientsMu.Lock()
	client.Conn = conn
	ClientsMu.Unlock()

	fyne.Do(func() {
		logs.Logs.Add(widget.NewLabel("Client connected on this IP address: " + ip))
	})

	broadcastUserNameList()
	defer func() {
		conn.Close()

		ClientsMu.Lock()
		if c, ok := Clients[ip]; ok && c.Conn == conn {
			c.Conn = nil
		}
		ClientsMu.Unlock()

		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("The connection was interrupted on this IP address: " + ip))
		})
		broadcastUserNameList()
	}()

	for {
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Client disconnected on this IP address: " + ip))
			})
			break
		}
	}
}

func broadcastUserNameList() {
	var names []string

	ClientsMu.RLock()
	for _, val := range Clients {
		if val.Conn != nil {
			names = append(names, val.Name)
		}
	}
	ClientsMu.RUnlock()

	payload := WSResponse{
		Type: "USER_NAME_LIST",
		Data: names,
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
