package components

import (
	"Shat/logs"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LogContainer() *container.Scroll {
	logs.Logs = container.NewVBox(
		widget.NewLabel("See the logs here."),
	)

	return container.NewScroll(logs.Logs)
}
