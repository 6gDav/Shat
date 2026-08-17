package history

import "sync"

type ChatMessage struct {
	Type     string `json:"type"`
	FromIp   string `json:"fromip"`
	Message  string `json:"message"`
	RoomId   string `json:"roomId"`
	UserName string `json:"username"`
}

var (
	ChatHistory = make(map[string][]ChatMessage)
	ChatStoreMu sync.RWMutex
)
