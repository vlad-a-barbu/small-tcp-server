package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConnection(clientId string) {
	for {
		serverTime := time.Now().Format(time.ANSIC)
		clients[clientId].Write([]byte(serverTime))

		clientReader := bufio.NewReader(clients[clientId])
		request, err := clientReader.ReadString('\n')

		switch err {
			case io.EOF: {
				fmt.Printf("\nClosing connection with %s\n", clientId)
				delete(clients, clientId)
				fmt.Printf("Num active connections: %d\n\n", len(clients))
				return
			}
			case nil:{
				handleRequest(clientId, request)
			}
			default: {
				fmt.Println(err)
				return
			}
		}
	}
}

func handleRequest(clientId string, request string) {
	fmt.Printf("\nReceived request from client %s: %s\n", clientId, request)

	response := fmt.Sprintf("Invalid request. %s\n")

	clients[clientId].Write([]byte(response))
}

var clients map[string]net.Conn

func createListener() net.Listener {
	if len(os.Args) != 2 {
		log.Fatalln("Invalid args. Please provide a port number.")
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

func main() {

	listener := createListener()
	defer listener.Close()

	clients = make(map[string]net.Conn)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		clientId := uuid.New().String()

		clients[clientId] = conn

		fmt.Printf("\nNew connection accepted - id: %s\n", clientId)
		fmt.Printf("Num active connections: %d\n", len(clients))

		go handleConnection(clientId)
	}
}
