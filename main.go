package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./hosted/build")))

	ipadress, err_ip := getIPAddress()

	if err_ip != nil {
		fmt.Println("Error occured ", err_ip)
	}

	concatedIp := ipadress + ":3000"

	fmt.Printf("Server is running on http://%s\n", concatedIp)
	err := http.ListenAndServe(concatedIp, nil)

	if err != nil {
		fmt.Println("Error occured while trying to start the server")
	}
}

func getIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no active non-loopback IPv4 address found")
}
