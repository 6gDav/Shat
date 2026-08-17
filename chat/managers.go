package chat

import (
	"hosting_login_page/history"
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

type ChatMessage struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	RoomID   string `json:"room_id"`
	UserName string `json:"username"`
}

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

func (cs *ManageChat) PrivateChat(targetIP string, text string, username string, saveChat bool) {
	roomId := generateRoomID(cs.IP, targetIP)

	//out
	outMsg := ChatMessage{
		Type:     "private",
		From:     cs.IP,
		To:       targetIP,
		Text:     text,
		RoomID:   roomId,
		UserName: username,
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

		if saveChat {
			history.ChatStoreMu.Lock()
			history.ChatHistory[roomId] = append(history.ChatHistory[roomId], history.ChatMessage{
				Ip:       cs.IP,
				Type:     "private",
				UserName: username,
				Message:  text,
				RoomId:   roomId,
			})
			history.ChatStoreMu.Unlock()
		}
	}

	if senderClient != nil && senderClient.Conn != nil {
		_ = senderClient.Conn.WriteJSON(outMsg)
	}
}

func (cs *ManageChat) GroupChat(text string, username string, saveChat bool) {

	ClientsMu.RLock()
	clientsCopy := make([]*Client, 0, len(Clients))
	for _, client := range Clients {
		if client != nil && client.Conn != nil {
			clientsCopy = append(clientsCopy, client)
		}
	}
	ClientsMu.RUnlock()

	if saveChat {
		history.ChatStoreMu.Lock()
		history.ChatHistory["Group"] = append(history.ChatHistory["Group"], history.ChatMessage{
			Ip:       cs.IP,
			Type:     "Group",
			UserName: username,
			Message:  text,
			RoomId:   "Group",
		})
		history.ChatStoreMu.Unlock()
	}

	for _, client := range clientsCopy {
		outMsg := ChatMessage{
			Type:     "group",
			From:     cs.IP,
			To:       client.IP,
			Text:     text,
			RoomID:   "Group",
			UserName: username,
		}

		err := client.Conn.WriteJSON(outMsg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Error occurred while sending group message: " + err.Error()))
			})
		}
	}
}

// func (cs *ManageChat) SendChatHistory() {
// 	history.ChatStoreMu.Lock()
// 	defer history.ChatStoreMu.Unlock()

// 	for roomIDKey, messagesValue := range history.ChatHistory {
// 		if strings.Contains(roomIDKey, cs.IP) {
// 			for _, msg := range messagesValue {
// 				outMsg := ChatMessage{
// 					Type:     "private_history",
// 					From:     msg.Ip,
// 					Text:     msg.Message,
// 					RoomID:   roomIDKey,
// 					UserName: msg.UserName,
// 				}
// 				_ = cs.Conn.WriteJSON(outMsg)
// 			}
// 		}
// 		if roomIDKey == "Group" {
// 			for _, msg := range messagesValue {
// 				outMsg := ChatMessage{
// 					Type:     "group_history",
// 					From:     msg.Ip,
// 					Text:     msg.Message,
// 					RoomID:   "Group",
// 					UserName: msg.UserName,
// 				}
// 				_ = cs.Conn.WriteJSON(outMsg)
// 			}
// 		}
// 	}
// }
