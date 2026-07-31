package components

import (
	"hosting_login_page/logs"
	runmanagger "hosting_login_page/runManagger"

	"fyne.io/fyne/v2/widget"
)

func StartButtonElement() *widget.Button {
	var clicked bool = false
	var button *widget.Button
	//initialText := "Start"

	button = widget.NewButton("Start", func() {
		if clicked {
			logs.Logs.Add(widget.NewLabel("Servers stopped..."))
			clicked = false
			button.SetText("Start")
			runmanagger.StopServer()
		} else {
			logs.Logs.Add(widget.NewLabel("Servers started..."))
			clicked = true
			button.SetText("Stop")
			runmanagger.StartServer()
		}
	})

	return button
}
