package server

import (
	"hosting_login_page/server/logic"
	"net/http"
)

var HttpServer *http.Server

func SetAPIendpoint() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", logic.ServePage)

	mux.HandleFunc("GET /redirectuser", logic.RedirectUser)

	mux.HandleFunc("GET /restorechathistory", logic.RestoreChatHistory)

	mux.HandleFunc("POST /submit", logic.SubmitUserName)

	mux.HandleFunc("PATCH /submitnewusername", logic.SubmitNewUserName)

	mux.HandleFunc("GET /setws", logic.WSHandsShake)

	portStart(mux)
}
