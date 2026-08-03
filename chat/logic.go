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
	var names []string

	ClientsMu.RLock()
	for _, val := range Clients {
		names = append(names, val.Name)
	}
	ClientsMu.RUnlock()

	namesData, err := json.Marshal(names)
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

	//managger
	ClientsMu.Lock()
	client, exists := Clients[ip]
	ClientsMu.Unlock()

	if !validateClientExistence(exists, conn) {
		return
	}
	manageConnection(client, conn, ip)
}

/*
err = conn.WriteMessage(websocket.TextMessage, []byte("Üdvözöllek a szerveren!"))
    if err != nil {
        fyne.Do(func() {
            logs.Logs.Add(widget.NewLabel("Üzenet küldése sikertelen: " + err.Error()))
        })
        conn.Close()
        return
    }
*/
