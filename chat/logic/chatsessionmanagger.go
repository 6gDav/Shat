package logic

import (
	"Shat/chat/client"

	"github.com/gorilla/websocket"
)

type ManageChat struct {
	IP     string
	Conn   *websocket.Conn
	Client *client.Client
}

type ChatMessage struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	RoomID   string `json:"room_id"`
	UserName string `json:"username"`
}

func NewClientSession(client *client.Client, conn *websocket.Conn, ip string) *ManageChat {
	return &ManageChat{
		IP:     ip,
		Conn:   conn,
		Client: client,
	}
}

func (cs *ManageChat) ManageConnection() {
	client.ClientsMu.Lock()
	cs.Client.Conn = cs.Conn
	client.ClientsMu.Unlock()

	BroadcastUserNameList()
	defer func() {
		cs.Conn.Close()

		client.ClientsMu.Lock()
		if c, ok := client.Clients[cs.IP]; ok && c.Conn == cs.Conn {
			c.Conn = nil
		}
		client.ClientsMu.Unlock()

		BroadcastUserNameList()
	}()
	cs.manageChat()
}

func (cs *ManageChat) manageChat() {
	for {
		//in
		var msg map[string]string
		err := cs.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		targetIP := msg["target_ip"]
		text := msg["text"]
		username := msg["userName"]

		messagetype := msg["type"]

		switch messagetype {
		case "private":
			cs.PrivateChat(targetIP, text, username, true)
		case "group":
			cs.GroupChat(text, username, true)
		}
	}
}
