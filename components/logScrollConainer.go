package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Logs *fyne.Container

func LogContainer() *container.Scroll {
	Logs = container.NewVBox(
		widget.NewLabel("See the logs here."),
	)

	return container.NewScroll(Logs)
}
