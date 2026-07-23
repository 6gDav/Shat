package main

import (
	"hosting_login_page/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("ChatApp")

	header := components.TitleElement()

	main_elements := container.NewVBox(
		widget.NewLabel("Intended port (3001, 8000, 8080 recommended)"),
		components.PortElement(),
		widget.NewLabel("Start Servers"),
		components.StartButtonElement(),
		widget.NewLabel("Create QR code"),
		components.QRCodeButtonElement(mainApp),
	)

	centeringElements := container.NewHBox(layout.NewSpacer(), main_elements, layout.NewSpacer())

	buttons := container.NewHBox(
		layout.NewSpacer(),
	)

	mainDesign := container.NewBorder(header, buttons, nil, nil, centeringElements)
	endDesign := container.NewPadded(mainDesign)

	mainWindow.SetContent(endDesign)
	mainWindow.Resize(fyne.NewSize(400, 350))
	mainWindow.ShowAndRun()
}
