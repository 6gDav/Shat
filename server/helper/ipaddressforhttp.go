package helper

import (
	"net"
	"net/http"
)

func GetIpAddressForEndPints(r *http.Request) (string, string, error) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	return ip, port, err
}
