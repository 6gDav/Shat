package main

import (
	_ "embed"

	"Shat/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//go:embed components/appIcon.png
var iconBytes []byte

func main() {
	mainApp := app.New()

	icon := fyne.NewStaticResource("appIcon.png", iconBytes)
	mainApp.SetIcon(icon)

	mainWindow := mainApp.NewWindow("Shat")

	header := components.TitleElement()

	controllElemets := container.NewVBox(
		widget.NewLabel("Start Servers"),
		components.StartButtonElement(),
		widget.NewLabel("Create QR code"),
		components.QRCodeButtonElement(mainApp),
	)

	buttons := container.NewHBox(
		layout.NewSpacer(),
		controllElemets,
		layout.NewSpacer(),
	)

	content := container.NewBorder(
		buttons,
		nil,
		nil,
		nil,
		components.LogContainer(),
	)

	mainDesign := container.NewBorder(header, nil, nil, nil, content)
	endDesign := container.NewPadded(mainDesign)

	mainWindow.SetContent(endDesign)
	mainWindow.Resize(fyne.NewSize(400, 450))
	mainWindow.ShowAndRun()
}
