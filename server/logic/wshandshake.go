package logic

import (
	"hosting_login_page/chat/logic"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandsShake(w http.ResponseWriter, r *http.Request) {
	logic.HandShake(w, r, &upgrader)
}
