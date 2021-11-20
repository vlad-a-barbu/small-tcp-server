package manager

import (
	"../handlers"
	"../utils"
	"fmt"
	"strings"
)

func handleRequest(clientId string, payload string) string {
	fmt.Printf("\nReceived request from client %s: %s\n", clientId, payload)

	tokens := utils.Tokenize(payload)
	command := tokens[0]

	var response string

	switch command {
		case "squares": {
			fmt.Println("Processing squares request")
			response = handlers.Squares(tokens)
			break
		}
		case "scramble": {
			fmt.Println("Processing scramble request")
			response = handlers.Scramble(tokens)
			break
		}
		default: {
			fmt.Println("Unable to process request")
			response = "Invalid request"
			break
		}
	}

	return strings.TrimSpace(response) + "\n"
}

