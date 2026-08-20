package logic

import (
	"encoding/json"
	"hosting_login_page/chat/client"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/http"

	"fyne.io/fyne/v2/widget"
)

func SubmitUserName(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := helper.GetIpAddressForEndPints(r)

	var data struct {
		Name string `json:"name"`
	}

	errDecode := json.NewDecoder(r.Body).Decode(&data)
	if errDecode != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	client.ClientsMu.Lock()
	if _, exists := client.Clients[ip]; !exists {
		client.Clients[ip] = &client.Client{
			IP:   ip,
			Name: data.Name,
			Conn: nil,
		}
	}
	client.ClientsMu.Unlock()

	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"ip": ip,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logs.Logs.Add(widget.NewLabel("Failed to send IP address to client: " + err.Error()))
	}
}
