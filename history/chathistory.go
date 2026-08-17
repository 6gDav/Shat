package history

import "sync"

type ChatMessage struct {
	Ip       string `json:"ip"`
	Type     string `json:"type"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

var (
	ChatHistory = make(map[string][]ChatMessage)
	ChatStoreMu sync.RWMutex
)
