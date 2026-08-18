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

func SubmitUserName(w http.ResponseWriter, r *http.Request) {
	//Ip adress fetch
	ip, _, _ := helper.GetIpAddressForEndPints(r)

	//Name fetch
	var data struct {
		Name string `json:"name"`
	}

	errdecode := json.NewDecoder(r.Body).Decode(&data)
	if errdecode != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	chat.ClientsMu.Lock()
	if _, exists := chat.Clients[ip]; !exists {
		chat.Clients[ip] = &chat.Client{
			IP:   ip,
			Name: data.Name,
			Conn: nil,
		}
		//Loging out
		logMsg := fmt.Sprintf("%+v\n", chat.Clients)
		logs.Logs.Add(widget.NewLabel("New user connected to the server:  " + logMsg))
	}
	chat.ClientsMu.Unlock()

	//Return 200 (Ok)
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"ip": ip,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logs.Logs.Add(widget.NewLabel("Failed to send ip andress to the client side: " + err.Error()))
	}
}
