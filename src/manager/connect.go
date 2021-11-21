package manager

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/small-tcp-server/v1/src/config"
	"io"
	"net"
	"time"
)

func Connect(){

	listener := createListener()
	defer listener.Close()

	clients := make(map[string]net.Conn)

	config := config.Read()

	for {

		if len(clients) == config.MaxConnections {
			continue
		}

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		clientId := uuid.New().String()
		clients[clientId] = conn

		fmt.Printf("\nNew client accepted - id: %s\n", clientId)
		fmt.Printf("Num active connections: %d\n", len(clients))

		serverTime := time.Now().Format(time.ANSIC)
		message := fmt.Sprintf("Connected successfully %s\n", serverTime)
		conn.Write([]byte(message))

		go handleConnection(clients, clientId, config)
	}
}

func handleConnection(clients map[string]net.Conn, clientId string, cfg config.Configuration){

	numRequests := 0

	for {
		conn := clients[clientId]
		clientReader := bufio.NewReader(conn)
		request, err := clientReader.ReadString('\n')

		numRequests += 1
		if numRequests > cfg.MaxRequests {
			fmt.Printf("\nClosing connection with %s\n", clientId)
			clients[clientId].Write([]byte("Max requests limit exceeded, closing connection\n"))
			delete(clients, clientId)
			fmt.Printf("Num active connections: %d\n\n", len(clients))
			return
		}

		switch err {
		case io.EOF:
			{
				fmt.Printf("\nClosing connection with %s\n", clientId)
				delete(clients, clientId)
				fmt.Printf("Num active connections: %d\n\n", len(clients))
				return
			}
		case nil:
			{
				response := handleRequest(clientId, request, cfg.MaxArgs)
				conn.Write([]byte(response))
			}
		default:
			{
				fmt.Println(err)
				return
			}
		}
	}
}

func removeClient(clients map[string]net.Conn, clientId string) {
	fmt.Printf("\nClosing connection with %s\n", clientId)
	clients[clientId].Write([]byte(""))
	delete(clients, clientId)
	fmt.Printf("Num active connections: %d\n\n", len(clients))
}
