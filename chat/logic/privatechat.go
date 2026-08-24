package logic

import (
	"Shat/chat/client"
	"Shat/chat/helper"
	"Shat/history"
	"Shat/logs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (cs *ManageChat) PrivateChat(targetIP string, text string, username string, saveChat bool) {
	roomId := helper.GenerateRoomID(cs.IP, targetIP)

	outMsg := ChatMessage{
		Type:     "private",
		From:     cs.IP,
		To:       targetIP,
		Text:     text,
		RoomID:   roomId,
		UserName: username,
	}

	client.ClientsMu.RLock()
	senderClient := client.Clients[cs.IP]
	targetClient := client.Clients[targetIP]
	client.ClientsMu.RUnlock()

	if targetClient != nil && targetClient.Conn != nil {
		err := targetClient.Conn.WriteJSON(outMsg)
		if err != nil {
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel("Error occurred while sending message: " + err.Error()))
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
