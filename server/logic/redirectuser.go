package logic

import (
	"encoding/json"
	"hosting_login_page/chat/client"
	"hosting_login_page/server/helper"
	"net/http"
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

	client.ClientsMu.Lock()
	user, exists := client.Clients[ip]
	if exists {
		response = RedirectResponse{
			Redirect:  true,
			UserName:  user.Name,
			IpAddress: user.IP,
		}
	} else {
		response = RedirectResponse{
			Redirect: false,
		}
	}
	client.ClientsMu.Unlock()

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
}
