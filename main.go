package main

import (
	"hosting_login_page/server"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
	server.SetAPIendpoint()
	server.SetMDNSserver()

}
