package logic

import "net/http"

var fileServer = http.FileServer(http.Dir("../page/build"))

func ServePage(w http.ResponseWriter, r *http.Request) {
	fileServer.ServeHTTP(w, r)
}
