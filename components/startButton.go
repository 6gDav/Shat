package components

import (
	"hosting_login_page/logs"
	"hosting_login_page/runmanager"

	"fyne.io/fyne/v2/widget"
)

func StartButtonElement() *widget.Button {
	var clicked bool = false
	var button *widget.Button

	button = widget.NewButton("Start", func() {
		if clicked {
			logs.Logs.Add(widget.NewLabel("Servers stopped..."))
			clicked = false
			button.SetText("Start")
			runmanager.StopServer()
		} else {
			logs.Logs.Add(widget.NewLabel("Servers started..."))
			clicked = true
			button.SetText("Stop")
			runmanager.StartServer()
		}
	})

	return button
}
