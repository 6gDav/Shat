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
	//Closing mDNS service
	components.ClosingmDNS()

	//Closing HTTP connections
	components.ClosingHTTP()
}
