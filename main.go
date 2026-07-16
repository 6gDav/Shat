package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/mdns"
)

const port = 3000

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./hosted/build"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	go func() {
		fmt.Printf("Server is running on port %d\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
		if err != nil {
			log.Fatalf("Error occurred while trying to start the server: %v", err)
		}
	}()

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

func getIPAddress() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, fmt.Errorf("no active non-loopback IPv4 address found")
}
