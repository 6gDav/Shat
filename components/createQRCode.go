package components

import (
	"fyne.io/fyne/v2/widget"
)

func QRCodeButtonElement() *widget.Button {
	button := widget.NewButton("Create QR code", func() {
		println("The button has pressed")
	})

	return button
}
