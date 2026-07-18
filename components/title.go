package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func TitleElement() *widget.Label {
	title := widget.NewLabelWithStyle("ChatApp", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return title
}
