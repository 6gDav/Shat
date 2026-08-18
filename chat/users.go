package chat

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	IP   string          `json:"ip"`
	Name string          `json:"name"`
	Conn *websocket.Conn `json:"-"`
}

func (c *Client) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{IP: %s, Name: %s}", c.IP, c.Name)
}

var (
	Clients   = make(map[string]*Client)
	ClientsMu sync.RWMutex
)
