package logic

import (
	"encoding/json"
	"hosting_login_page/chat"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/http"

	"fyne.io/fyne/v2/widget"
)

func RedirectUser(w http.ResponseWriter, r *http.Request) {
	ip, _, err := helper.GetIpAddressForEndPints(r)
	if err != nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
	}

	type RedirectResponse struct {
		Redirect  bool   `json:"redirect"`
		UserName  string `json:"userName"`
		IpAddress string `json:"ipAddress"`
	}

	var response RedirectResponse

	chat.ClientsMu.Lock()
	client, exists := chat.Clients[ip]
	if exists {
		response = RedirectResponse{
			Redirect:  true,
			UserName:  client.Name,
			IpAddress: client.IP,
		}
	} else {
		response = RedirectResponse{
			Redirect: false,
		}
	}
	chat.ClientsMu.Unlock()

	if response.Redirect {
		logs.Logs.Add(widget.NewLabel("User redirected to dashboard IP: " + ip))
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
}
