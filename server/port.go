package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

var Port int = 3000

func PortStart(muxInstance *http.ServeMux) {
	go func() {
		for i := 0; i < 10; i++ {
			currentPort := Port + i

			HttpServer = &http.Server{
				Addr:    fmt.Sprintf(":%d", currentPort),
				Handler: muxInstance,
			}

			log.Printf("Trying to start the server %d \n", currentPort)

			err := HttpServer.ListenAndServe()

			if err != nil && isPortInUse(err) {
				log.Printf("This port is occupied %d ", currentPort)
				continue
			}

			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Http server is stoped")
				break
			}

			if err != nil {
				log.Printf("Server error: %v\n", err)
				break
			}
		}
	}()
}

func isPortInUse(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return opErr.Op == "listen"
	}
	return false
}
