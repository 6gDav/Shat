package qrcode

import (
	"hosting_login_page/server"

	"github.com/skip2/go-qrcode"
)

func CreateQRCode() ([]byte, error) {
	image, err := qrcode.Encode(server.URL, qrcode.Medium, 256)

	return image, err
}
