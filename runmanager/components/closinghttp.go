package components

import (
	"Shat/chat/client"
	"Shat/logs"
	"Shat/server"
	"context"
	"fmt"
	"time"

	"fyne.io/fyne/v2/widget"
)

func ClosingHTTP() {
	if server.HttpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := server.HttpServer.Shutdown(ctx); err != nil {
			errMsg := fmt.Sprintf("Error occurred while trying to stop the server: %v", err)
			logs.Logs.Add(widget.NewLabel(errMsg))
		} else {
			logs.Logs.Add(widget.NewLabel("HTTP server shut down successfully"))
		}
		server.HttpServer = nil
	}

	client.ClientsMu.Lock()
	clear(client.Clients)
	client.ClientsMu.Unlock()
}
