package server

import (
	"Shat/server/logic"
	"net/http"
)

var HttpServer *http.Server

func SetAPIendpoint() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", logic.ServePage)

	mux.HandleFunc("GET /user/redirect", logic.RedirectUser)

	mux.HandleFunc("GET /chats/history", logic.RestoreChatHistory)

	mux.HandleFunc("POST /user/username", logic.SubmitUserName)

	mux.HandleFunc("PATCH /user/new/name", logic.SubmitNewUserName)

	mux.HandleFunc("GET /ws", logic.WSHandsShake)

	portStart(mux)
}
