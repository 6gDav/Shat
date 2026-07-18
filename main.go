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
	myApp := app.New()
	myWindow := myApp.NewWindow("ChatApp")

	header := components.TitleElement()

	main_elemts := container.NewVBox(
		widget.NewLabel("Start Servers"),
		components.StartButtonElement(),
		layout.NewSpacer(),
		widget.NewLabel("Create QR code"),
		components.QRCodeButtonElement(),
	)

	midofpage := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 140)), main_elemts)

	centeringElements := container.NewHBox(layout.NewSpacer(), midofpage, layout.NewSpacer())

	buttons := container.NewHBox(
		layout.NewSpacer(),
	)

	mainDesign := container.NewBorder(header, buttons, nil, nil, centeringElements)

	endDesign := container.NewPadded(mainDesign)

	myWindow.SetContent(endDesign)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()

}
