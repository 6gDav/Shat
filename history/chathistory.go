package history

import "sync"

type ChatMessage struct {
	Ip      string
	Message string
}

var (
	ChatHistory = make(map[string][]ChatMessage)
	ChatStoreMu sync.RWMutex
)
