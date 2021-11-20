package manager

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net"
	"time"
)

func Connect(){

	listener := createListener()
	defer listener.Close()

	clients := make(map[string]net.Conn)

	for {
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

		go handleConnection(clients, clientId)
	}
}

func handleConnection(clients map[string]net.Conn, clientId string){
	for {
		conn := clients[clientId]
		clientReader := bufio.NewReader(conn)
		request, err := clientReader.ReadString('\n')

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
				response := handleRequest(clientId, request)
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
