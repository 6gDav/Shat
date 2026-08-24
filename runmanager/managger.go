package runmanager

import (
	"Shat/runmanager/components"
	"Shat/server"
)

func StartServer() {
	server.SetAPIendpoint()
	server.SetMDNSserver()
}

func StopServer() {
	components.ClosingmDNS()
	components.ClosingHTTP()
}
