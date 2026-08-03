package chat

import (
	"encoding/json"
	"hosting_login_page/logs"

	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type WSResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func HandShake(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occurred while trying to set the Upgrader: " + err.Error()))
		})
		return
	}
	//Return names
	sendUserNameList(conn)

	//managger
	ClientsMu.Lock()
	client, exists := Clients[ip]
	ClientsMu.Unlock()

	if !validateClientExistence(exists, conn) {
		return
	}
	manageConnection(client, conn, ip)
}

func sendUserNameList(conn *websocket.Conn) {
	var names []string

	ClientsMu.RLock()
	for _, val := range Clients {
		names = append(names, val.Name)
	}
	ClientsMu.RUnlock()

	payload := WSResponse{
		Type: "USER_NAME_LIST",
		Data: names,
	}

	namesData, err := json.Marshal(payload)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occred while tryng to convert to JSON " + err.Error()))
		})
		errMsg := []byte("ERROR: Failed to serialize name list")
		_ = conn.WriteMessage(websocket.TextMessage, errMsg)

		return
	}

	err = conn.WriteMessage(websocket.TextMessage, namesData)
	if err != nil {
		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("Error occured while trying to send the name list: " + err.Error()))
		})
		conn.Close()
		return
	}
}
