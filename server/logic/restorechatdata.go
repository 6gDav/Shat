package logic

import (
	"encoding/json"
	"hosting_login_page/history"
	"hosting_login_page/server/helper"
	"net/http"
	"strings"
)

func RestoreChatHistory(w http.ResponseWriter, r *http.Request) {
	ip, _, err := helper.GetIpAddressForEndPints(r)
	if err != nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
	}

	var historyList []history.ChatMessage

	history.ChatStoreMu.Lock()
	defer history.ChatStoreMu.Unlock()

	for ipAddressKey, message := range history.ChatHistory {
		if strings.Contains(ipAddressKey, ip) {
			historyList = append(historyList, message...)
		} else if ipAddressKey == "Group" {
			historyList = append(historyList, message...)
		}
	}
	if err := json.NewEncoder(w).Encode(historyList); err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
}
