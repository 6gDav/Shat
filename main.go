package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/mdns"
)

// Segédfunkció a gép helyi IP címének lekérdezésére
func getLocalIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// Ellenőrizzük, hogy nem loopback-e (127.0.0.1) és hogy IP4-e
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

func main() {
	port := 3000

	// 1. Gépnév lekérése és tisztítása (.local duplázódás elkerülése)
	host, err := os.Hostname()
	if err != nil {
		host = "goserver"
	}
	// Ha a gépnév eleve ".local"-ra végződik, vágjuk le a biztonság kedvéért
	host = strings.TrimSuffix(host, ".local")
	host = strings.TrimSuffix(host, ".local.")

	// 2. Saját IP megkeresése
	localIP := getLocalIP()
	var ips []net.IP
	if localIP != nil {
		ips = []net.IP{localIP}
		fmt.Printf("Saját IP-cím észlelve: %s\n", localIP.String())
	} else {
		fmt.Println("Figyelem: Nem sikerült automatikusan meghatározni a helyi IP-címet!")
	}

	// 3. mDNS Szolgáltatás konfigurálása explicit IP-vel
	service, err := mdns.NewMDNSService(
		"goserver",     // Ezen a néven hirdetünk (goserver.local)
		"_http._tcp",   // Szolgáltatás típusa
		"",             // Domain (alapértelmezett "local.")
		host+".local.", // Host (most már garantáltan csak egyszer van a végén .local)
		port,           // Port
		ips,            // <--- Itt átadjuk az észlelt IP-nket, így nem fog hibára futni!
		[]string{"path=/"},
	)
	if err != nil {
		fmt.Printf("Hiba az mDNS konfiguráció során: %v\n", err)
		return
	}

	// mDNS Szerver indítása
	mdnsServer, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		fmt.Printf("Hiba az mDNS indításakor: %v\n", err)
		return
	}
	defer mdnsServer.Shutdown()

	// 4. HTTP Szerver indítása
	http.Handle("/", http.FileServer(http.Dir("./hosted/build")))

	fmt.Printf("\nSikeresen elindult!\n")
	fmt.Printf("Próbáld megnyitni a hálózaton: http://goserver.local:%d\n", port)
	if localIP != nil {
		fmt.Printf("Vagy IP alapján: http://%s:%d\n", localIP.String(), port)
	}

	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Println("Hiba történt a HTTP szerver futása közben:", err)
	}
}
