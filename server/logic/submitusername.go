package logic

import (
	"encoding/json"
	"hosting_login_page/chat/client"
	"hosting_login_page/server/helper"
	"net/http"
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

	client.ClientsMu.Lock()
	if _, exists := client.Clients[ip]; !exists {
		client.Clients[ip] = &client.Client{
			IP:   ip,
			Name: data.Name,
			Conn: nil,
		}
	}
	client.ClientsMu.Unlock()

	//Return 200 (Ok)
	w.WriteHeader(http.StatusOK)
}
