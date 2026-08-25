package logic

import (
	"io/fs"
	"net/http"
)

var fileServer http.Handler

func SetStaticAssets(embedFS fs.FS) {
	subFS, err := fs.Sub(embedFS, "page/build")
	if err != nil {
		panic("Falied to load the build folder: " + err.Error())
	}
	fileServer = http.FileServer(http.FS(subFS))
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	if fileServer != nil {
		fileServer.ServeHTTP(w, r)
	} else {
		http.Error(w, "Assets didnt loaded: ", http.StatusInternalServerError)
	}
}
