package server

import (
	"errors"
	"fmt"
	"hosting_login_page/logs"
	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Port int = 3000

func portStart(muxInstance *http.ServeMux) {
	go func() {
		for i := 0; i < 10; i++ {
			currentPort := Port + i

			HttpServer = &http.Server{
				Addr:    fmt.Sprintf(":%d", currentPort),
				Handler: muxInstance,
			}

			msg := fmt.Sprintf("Trying to start the server on this port: %d \n", currentPort)
			fyne.Do(func() {
				logs.Logs.Add(widget.NewLabel(msg))
			})

			err := HttpServer.ListenAndServe()

			if err != nil && isPortInUse(err) {
				msg := fmt.Sprintf("This port is occupied %d ", currentPort)
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

func isPortInUse(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return opErr.Op == "listen"
	}
	return false
}
