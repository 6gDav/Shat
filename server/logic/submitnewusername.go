package logic

import (
	"encoding/json"
	"fmt"
	"hosting_login_page/chat"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/http"

	"fyne.io/fyne/v2/widget"
)

func SubmitNewUserName(w http.ResponseWriter, r *http.Request) {
	//Ip adress fetch
	ip, _, err := helper.GetIpAddressForEndPints(r)

	if err != nil {
		logMsg := fmt.Sprintf("Error occurred while this user %s tried to change username: %v", ip, err)
		logs.Logs.Add(widget.NewLabel(logMsg))
	}

	//Name fetch
	var data struct {
		Name string `json:"name"`
	}

	defer func() {
		r.Body.Close()
		chat.BroadcastUserNameList()
	}()

	errdecode := json.NewDecoder(r.Body).Decode(&data)
	if errdecode != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	chat.ClientsMu.Lock()
	if client, exists := chat.Clients[ip]; exists {
		client.Name = data.Name
	}
	chat.ClientsMu.Unlock()

	//Return 200 (Ok)
	w.WriteHeader(http.StatusOK)
}
