package server

import (
	"fmt"
	"hosting_login_page/logs"

	"net"

	"fyne.io/fyne/v2/widget"
	"github.com/hashicorp/mdns"
)

var MdnsServer *mdns.Server
var domainName string = "loginpage"

func SetMDNSserver() {
	ipAddress, err := getIPAddress()
	if err != nil {
		logs.Logs.Add(widget.NewLabel("Error occurred while trying to fetch the IP address: " + err.Error()))
	}

	service, err := mdns.NewMDNSService(
		domainName,
		"_http._tcp",
		"local.",
		"loginpage.local.",
		Port,
		[]net.IP{ipAddress},
		[]string{"txtv=1"},
	)
	if err != nil {
		logs.Logs.Add(widget.NewLabel("Establish mDNS service is unsuccessful: " + err.Error()))
	}

	var errServer error
	MdnsServer, errServer = mdns.NewServer(&mdns.Config{Zone: service})
	if errServer != nil {
		logs.Logs.Add(widget.NewLabel("Error occurred while trying to start the mDNS server: " + err.Error()))
	}

	//fmt.Printf("Web page is here: http://loginpage.local:%d\n", Port)
	msg := fmt.Sprintf("Web page is available on this link: http://%s.local:%d", domainName, Port)
	logs.Logs.Add(widget.NewLabel(msg))

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan
	// fmt.Println("\nShutting down server...")
}
