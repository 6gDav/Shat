package chat

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	IP   string
	Name string
	Conn *websocket.Conn
}

var (
	Clients   = make(map[string]*Client)
	ClientsMu sync.RWMutex
)
