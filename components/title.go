package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func TitleElement() *widget.Label {
	title := widget.NewLabelWithStyle("Shat", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return title
}
