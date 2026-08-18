package logic

import (
	"hosting_login_page/chat"
	"hosting_login_page/history"
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (cs *ManageChat) GroupChat(text string, username string, saveChat bool) {

	chat.ClientsMu.RLock()
	clientsCopy := make([]*chat.Client, 0, len(chat.Clients))
	for _, client := range chat.Clients {
		if client != nil && client.Conn != nil {
			clientsCopy = append(clientsCopy, client)
		}
	}
	chat.ClientsMu.RUnlock()

	if saveChat {
		history.ChatStoreMu.Lock()
		history.ChatHistory["Group"] = append(history.ChatHistory["Group"], history.ChatMessage{
			Type:     "group",
			FromIp:   cs.IP,
			Message:  text,
			RoomId:   "Group",
			UserName: username,
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
