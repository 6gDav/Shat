package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/mdns"
)

func setMDNSserver() {
	ipAddress, err := getIPAddress()
	if err != nil {
		log.Fatalf("Error occurred while trying to fetch IP address: %v", err)
	}

	service, err := mdns.NewMDNSService(
		"loginpage",
		"_http._tcp",
		"local.",
		"loginpage.local.",
		port,
		[]net.IP{ipAddress},
		[]string{"txtv=1"},
	)
	if err != nil {
		log.Fatalf("Establish mDNS service is unsuccessful: %v", err)
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Fatalf("Error occurred while trying to start the mDNS server: %v", err)
	}
	defer server.Shutdown()

	fmt.Printf("Web page is here: http://loginpage.local:%d\n", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("\nShutting down server...")
}
