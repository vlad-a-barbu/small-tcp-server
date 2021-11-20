package manager

import (
	"fmt"
	"log"
	"net"
	"os"
)

func createListener() net.Listener {
	if len(os.Args) != 2 {
		log.Fatalln("Provide a port number")
	}

	port := os.Args[1]
	address := fmt.Sprintf("localhost:%s", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Listening on port %s\n", port)

	return listener
}

