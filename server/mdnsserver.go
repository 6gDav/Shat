package server

import (
	"fmt"

	"log"
	"net"

	"github.com/hashicorp/mdns"
)

var MdnsServer *mdns.Server

func SetMDNSserver() {
	ipAddress, err := getIPAddress()
	if err != nil {
		log.Fatalf("Error occurred while trying to fetch IP address: %v", err)
	}

	service, err := mdns.NewMDNSService(
		"loginpage",
		"_http._tcp",
		"local.",
		"loginpage.local.",
		Port,
		[]net.IP{ipAddress},
		[]string{"txtv=1"},
	)
	if err != nil {
		log.Fatalf("Establish mDNS service is unsuccessful: %v", err)
	}

	var errServer error
	MdnsServer, errServer = mdns.NewServer(&mdns.Config{Zone: service})
	if errServer != nil {
		log.Fatalf("Error occurred while trying to start the mDNS server: %v", errServer)
	}

	fmt.Printf("Web page is here: http://loginpage.local:%d\n", Port)

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan
	// fmt.Println("\nShutting down server...")
}
