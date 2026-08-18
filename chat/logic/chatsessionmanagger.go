package logic

import (
	"hosting_login_page/chat"

	"github.com/gorilla/websocket"
)

type ManageChat struct {
	IP     string
	Conn   *websocket.Conn
	Client *chat.Client
}

type ChatMessage struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	RoomID   string `json:"room_id"`
	UserName string `json:"username"`
}

func NewClientSession(client *chat.Client, conn *websocket.Conn, ip string) *ManageChat {
	return &ManageChat{
		IP:     ip,
		Conn:   conn,
		Client: client,
	}
}

func (cs *ManageChat) ManageConnection() {
	chat.ClientsMu.Lock()
	cs.Client.Conn = cs.Conn
	chat.ClientsMu.Unlock()

	chat.BroadcastUserNameList()
	defer func() {
		cs.Conn.Close()

		chat.ClientsMu.Lock()
		if c, ok := chat.Clients[cs.IP]; ok && c.Conn == cs.Conn {
			c.Conn = nil
		}
		chat.ClientsMu.Unlock()

		chat.BroadcastUserNameList()
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
