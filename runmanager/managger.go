package runmanager

import (
	"hosting_login_page/runmanager/components"
	"hosting_login_page/server"
)

func StartServer() {
	server.SetAPIendpoint()
	server.SetMDNSserver()
}

func StopServer() {
	components.ClosingmDNS()
	components.ClosingHTTP()
}
