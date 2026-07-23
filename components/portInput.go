package components

import "fyne.io/fyne/v2/widget"

var portInput *widget.Entry

func PortElement() *widget.Entry {
	portInput = widget.NewEntry()
	portInput.SetPlaceHolder("Write the port here")

	return portInput
}
