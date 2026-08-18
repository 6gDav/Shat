package server

import (
	"errors"
	"fmt"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var port int = 3000

func portStart(muxInstance *http.ServeMux) {
	port := 3000
	go func() {
		for i := 0; i < 10; i++ {
			port = port + i

			HttpServer = &http.Server{
				Addr:    fmt.Sprintf(":%d", port),
				Handler: muxInstance,
			}

			msg := fmt.Sprintf("Trying to start the server on this port: %d \n", port)
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel(msg))
			})

			err := HttpServer.ListenAndServe()

			if err != nil && helper.IsPortInUse(err) {
				msg := fmt.Sprintf("This port is occupied %d ", port)
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel(msg))
				})
				continue
			}

			if errors.Is(err, http.ErrServerClosed) {
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel("HTTP server shot down"))
				})
				break
			}

			if err != nil {
				fyne.Do(func() {
					logs.Logs.Add(widget.NewLabel("Error occured: " + err.Error()))
				})
				break
			}
		}
	}()
}
