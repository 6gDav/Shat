package components

import (
	"bytes"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
)

var link string = "http://loginpage.local:3000"

func QRCodeButtonElement(fyneApp fyne.App) *widget.Button {
	button := widget.NewButton("Create QR code", func() {

		image, err := qrcode.Encode(link, qrcode.Medium, 256)
		if err != nil {
			log.Fatalf("Error occured while trying to create QR code: %v", err)
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
