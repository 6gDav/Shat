package components

import (
	"hosting_login_page/server"

	"fyne.io/fyne/v2/widget"
)

func StartButtonElement() *widget.Button {
	button := widget.NewButton("Start", func() {
		server.SetAPIendpoint()
		server.SetMDNSserver()
	})

	return button
}
