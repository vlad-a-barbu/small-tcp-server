package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnection(clientId string) {
	for {
		clientReader := bufio.NewReader(clients[clientId])
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
				clients[clientId].Write([]byte(response))
			}
		default:
			{
				fmt.Println(err)
				return
			}
		}
	}
}

func tokenize(request string) []string {

	tokens := strings.Split(strings.TrimSpace(request), " ")

	validTokens := []string{}
	for _, token := range tokens {
		if len(token) > 0 {
			validTokens = append(validTokens, token)
		}
	}

	return validTokens
}

func parseNum(token string) int {

	var num string
	for _, char := range token {
		if char >= '0' && char <= '9' {
			num += string(char)
		}
	}

	if len(num) > 0 {
		result, _ := strconv.Atoi(num)
		return result
	}

	return -1
}

func isSquare(num int) bool {
	return math.Sqrt(float64(num)) == math.Floor(math.Sqrt(float64(num)))
}

func getSquares(nums []int) []int {
	squares := []int{}
	for _, num := range nums {
		if (isSquare(num)){
			squares = append(squares, num)
		}
	}
	return squares
}

func handleRequest(clientId string, request string) string {
	fmt.Printf("\nReceived request from client %s: %s\n", clientId, request)

	tokens := tokenize(request)
	command := tokens[0]

	var response string

	switch command {
		case "squares": {
			if len(tokens) < 2 {
				response = fmt.Sprintf("Invalid args. Provide at least one token.\n")
				break
			} else {
				nums := []int{}
				for _, token := range tokens[1:] {
					num := parseNum(token)
					if num != -1 {
						nums = append(nums, num)
					}
				}
				squares := getSquares(nums)
				if len(squares) > 0 {
					response = "Square nums found: "
					for _, square := range(squares) {
						response += strconv.Itoa(square) + " "
					}
				} else {
					response = "No squares were found"
				}
			}
			response = strings.TrimSpace(response) + "\n"
			break
		}
		case "scramble": {
			if len(tokens) < 3 {
				response = fmt.Sprintf("Invalid args. Provide at least 2 tokens.\n")
				break
			} else {
				length := len(tokens[1])
				for i := 0; i < length; i++ {
					var scrambledToken string
					for _, token := range tokens[1:] {
						if length != len(token){
							response = fmt.Sprintf("Invalid args. Token lengths must be equal.\n")
							break
						}
						scrambledToken += string(token[i])
					}
					response += scrambledToken + " "
				}
			}
			response = strings.TrimSpace(response) + "\n"
			break
		}
		default: {
			response = fmt.Sprintf("Invalid request.\n")
			break
		}
	}

	return response
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

		serverTime := time.Now().Format(time.ANSIC)
		message := fmt.Sprintf("Connected successfully %s\n", serverTime)
		conn.Write([]byte(message))

		go handleConnection(clientId)
	}
}
