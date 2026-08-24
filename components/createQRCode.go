package components

import (
	"Shat/logs"
	"Shat/qrcode"
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func QRCodeButtonElement(fyneApp fyne.App) *widget.Button {
	button := widget.NewButton("Create QR code", func() {

		image, err := qrcode.CreateQRCode()
		if err != nil {
			logs.Logs.Add(widget.NewLabel("Failed to generate QR code: " + err.Error()))
		}

		qrCodeWindow := fyneApp.NewWindow("QR Code")
		img := canvas.NewImageFromReader(bytes.NewReader(image), "qrcode.png")
		img.SetMinSize(fyne.NewSize(256, 256))

		qrCodeWindow.SetContent(container.NewCenter(img))
		qrCodeWindow.Resize(fyne.NewSize(300, 300))
		qrCodeWindow.Show()
	})

	return button
}
