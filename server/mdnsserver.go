package server

import (
	"fmt"
	"hosting_login_page/logs"
	"hosting_login_page/server/helper"
	"net/url"

	"net"

	"fyne.io/fyne/v2/widget"
	"github.com/hashicorp/mdns"
)

var MdnsServer *mdns.Server
var URL string

func SetMDNSserver() {
	domainName := "shat"

	ipAddress, err := helper.GetIPAddressFormDNS()
	if err != nil {
		logs.Logs.Add(widget.NewLabel("Failed to fetch IP address: " + err.Error()))
	}

	fullDomainName := domainName + ".local"

	service, err := mdns.NewMDNSService(
		domainName,
		"_http._tcp",
		"local.",
		fullDomainName+".",
		port,
		[]net.IP{ipAddress},
		[]string{"txtv=1"},
	)
	if err != nil {
		logs.Logs.Add(widget.NewLabel("Failed to establish DNS service: %v " + err.Error()))
	}

	var errServer error
	MdnsServer, errServer = mdns.NewServer(&mdns.Config{Zone: service})
	if errServer != nil {
		logs.Logs.Add(widget.NewLabel("Error occurred while trying to start the DNS server: " + err.Error()))
	}

	vuildedURL := fmt.Sprintf("http://%s.local:%d", domainName, port)
	URL = vuildedURL
	logs.Logs.Add(widget.NewLabel("Web page is available at this link: " + vuildedURL))

	parsedURL, _ := url.Parse(vuildedURL)
	logs.Logs.Add(widget.NewHyperlink("Click here to open the web page!", parsedURL))
}
