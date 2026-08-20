package logic

import (
	"encoding/json"
	"fmt"
	"hosting_login_page/chat/client"
	"hosting_login_page/chat/logic"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/http"

	"fyne.io/fyne/v2/widget"
)

func SubmitNewUserName(w http.ResponseWriter, r *http.Request) {
	ip, _, err := helper.GetIpAddressForEndPints(r)

	if err != nil {
		logMsg := fmt.Sprintf("Error occurred while this user %s tried to change username: %v", ip, err)
		logs.Logs.Add(widget.NewLabel(logMsg))
	}

	var data struct {
		Name string `json:"name"`
	}

	defer func() {
		r.Body.Close()
		logic.BroadcastUserNameList()
	}()

	errdecode := json.NewDecoder(r.Body).Decode(&data)
	if errdecode != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	client.ClientsMu.Lock()
	if client, exists := client.Clients[ip]; exists {
		client.Name = data.Name
	}
	client.ClientsMu.Unlock()

	w.WriteHeader(http.StatusOK)
}
