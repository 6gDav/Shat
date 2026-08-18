package logic

import (
	"hosting_login_page/chat"
	"hosting_login_page/chat/helper"
	"hosting_login_page/history"
	"hosting_login_page/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (cs *ManageChat) PrivateChat(targetIP string, text string, username string, saveChat bool) {
	roomId := helper.GenerateRoomID(cs.IP, targetIP)

	//out
	outMsg := ChatMessage{
		Type:     "private",
		From:     cs.IP,
		To:       targetIP,
		Text:     text,
		RoomID:   roomId,
		UserName: username,
	}

	chat.ClientsMu.RLock()
	senderClient := chat.Clients[cs.IP]
	targetClient := chat.Clients[targetIP]
	chat.ClientsMu.RUnlock()

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
				Type:     "private",
				FromIp:   cs.IP,
				Message:  text,
				RoomId:   roomId,
				UserName: username,
			})
			history.ChatStoreMu.Unlock()
		}
	}

	if senderClient != nil && senderClient.Conn != nil {
		_ = senderClient.Conn.WriteJSON(outMsg)
	}
}
