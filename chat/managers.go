package chat

import (
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type ManageChat struct {
	IP     string
	Conn   *websocket.Conn
	Client *Client
}

type ReturnChatMessage struct {
	type string
	from string
	to   string
	text string
	room_id string
	userName string
}

/*
	"type":     "private",
	"from":     cs.IP,
	"to":       targetIP,
	"text":     text,
	"room_id":  generateRoomID(cs.IP, targetIP),
	"username": username,
*/

func NewClientSession(client *Client, conn *websocket.Conn, ip string) *ManageChat {
	return &ManageChat{
		IP:     ip,
		Conn:   conn,
		Client: client,
	}
}

func (cs *ManageChat) ManageConnection() {
	ClientsMu.Lock()
	cs.Client.Conn = cs.Conn
	ClientsMu.Unlock()

	fyne.Do(func() {
		logs.Logs.Add(widget.NewLabel("Client connected on this IP address: " + cs.IP))
	})

	BroadcastUserNameList()
	defer func() {
		cs.Conn.Close()

		ClientsMu.Lock()
		if c, ok := Clients[cs.IP]; ok && c.Conn == cs.Conn {
			c.Conn = nil
		}
		ClientsMu.Unlock()

		fyne.Do(func() {
			logs.Logs.Add(widget.NewLabel("The connection was interrupted on this IP address: " + cs.IP))
			BroadcastUserNameList()
		})
	}()
	cs.manageChat()
}

func (cs *ManageChat) manageChat() {
	defer cs.Conn.Close()

	for {
		//in
		var msg map[string]string
		err := cs.Conn.ReadJSON(&msg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Client disconnected on this IP address: " + cs.IP))
			})
			break
		}

		messagetype := msg["type"]

		switch messagetype {
		case "private":
			cs.privateChat(msg)
		case "group":
			cs.groupChat(msg)
		}
	}
}

func (cs ManageChat) privateChat(msg map[string]string) {
	targetIP := msg["target_ip"]
	text := msg["text"]
	username := msg["userName"]

	//out
	outMsg := map[string]string{
		"type":     "private",
		"from":     cs.IP,
		"to":       targetIP,
		"text":     text,
		"room_id":  generateRoomID(cs.IP, targetIP),
		"username": username,
	}

	ClientsMu.RLock()
	senderClient := Clients[cs.IP]
	targetClient := Clients[targetIP]
	ClientsMu.RUnlock()

	if targetClient != nil && targetClient.Conn != nil {
		err := targetClient.Conn.WriteJSON(outMsg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Error occurred while trying to send message: " + err.Error()))
			})
		}
	}

	if senderClient != nil && senderClient.Conn != nil {
		_ = senderClient.Conn.WriteJSON(outMsg)
	}
}

func (cs ManageChat) groupChat(msg map[string]string) {
	text := msg["text"]
	username := msg["userName"]

	ClientsMu.RLock()
	clientsCopy := make([]*Client, 0, len(Clients))
	for _, client := range Clients {
		if client != nil && client.Conn != nil {
			clientsCopy = append(clientsCopy, client)
		}
	}
	ClientsMu.RUnlock()

	for _, client := range clientsCopy {
		outMsg := map[string]string{
			"type":     "group",
			"from":     cs.IP,
			"to":       client.IP,
			"text":     text,
			"room_id":  "Group",
			"username": username,
		}

		err := client.Conn.WriteJSON(outMsg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Error occurred while sending group message: " + err.Error()))
			})
		}
	}
}
