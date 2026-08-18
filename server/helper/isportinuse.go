package helper

import (
	"errors"
	"net"
)

func IsPortInUse(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return opErr.Op == "listen"
	}
	return false
}
