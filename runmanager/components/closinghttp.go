package components

import (
	"context"
	"fmt"
	"hosting_login_page/chat/client"
	"hosting_login_page/logs"
	"hosting_login_page/server"
	"time"

	"fyne.io/fyne/v2/widget"
)

func ClosingHTTP() {
	if server.HttpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := server.HttpServer.Shutdown(ctx); err != nil {
			errMsg := fmt.Sprintf("Error occured while trying to stop the server: %v", err)
			logs.Logs.Add(widget.NewLabel(errMsg))
		} else {
			logs.Logs.Add(widget.NewLabel("HTTP serves succesfully shot down..."))
		}
		server.HttpServer = nil
	}

	client.ClientsMu.Lock()
	clear(client.Clients)
	client.ClientsMu.Unlock()
}
