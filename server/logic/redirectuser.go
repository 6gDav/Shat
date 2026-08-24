package logic

import (
	"Shat/chat/client"
	"Shat/server/helper"
	"encoding/json"
	"net/http"
)

type RedirectResponse struct {
	Redirect  bool   `json:"redirect"`
	UserName  string `json:"userName"`
	IpAddress string `json:"ipAddress"`
}

func RedirectUser(w http.ResponseWriter, r *http.Request) {
	ip, _, err := helper.GetIpAddressForEndPints(r)
	if err != nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
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
