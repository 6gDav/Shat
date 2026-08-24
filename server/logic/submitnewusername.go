package logic

import (
	"Shat/chat/client"
	"Shat/chat/logic"
	"Shat/logs"
	"Shat/server/helper"
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2/widget"
)

func SubmitNewUserName(w http.ResponseWriter, r *http.Request) {
	ip, _, err := helper.GetIpAddressForEndPints(r)

	if err != nil {
		logMsg := fmt.Sprintf("Failed to change username for %s: %v.", ip, err)
		logs.Logs.Add(widget.NewLabel(logMsg))
	}

	var data struct {
		Name string `json:"name"`
	}

	defer func() {
		r.Body.Close()
		logic.BroadcastUserNameList()
	}()

	errDecode := json.NewDecoder(r.Body).Decode(&data)
	if errDecode != nil {
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
