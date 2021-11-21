package manager

import (
	"fmt"
	"github.com/small-tcp-server/v1/src/handlers"
	"github.com/small-tcp-server/v1/src/utils"
	"strings"
)

func handleRequest(clientId string, payload string, maxArgs int) string {
	fmt.Printf("\nReceived request from client %s: %s\n", clientId, payload)

	tokens := utils.Tokenize(payload)

	if len(tokens) > maxArgs {
		return "Max args limit exceeded\n"
	}

	command := tokens[0]

	var response string

	switch command {
		case "squares": {
			fmt.Println("Processing squares request")
			response = handlers.Squares(tokens[1:])
			break
		}
		case "scramble": {
			fmt.Println("Processing scramble request")
			response = handlers.Scramble(tokens[1:])
			break
		}
		case "rsum": {
			fmt.Println("Processing rsum request")
			response = handlers.ReversedSum(tokens[1:])
			break
		}
		case "mean": {
			fmt.Println("Processing mean request")
			response = handlers.ArithmeticMean(tokens[1:])
			break
		}
		case "validpass": {
			fmt.Println("Processing validpass request")
			response = handlers.ValidPassword(tokens[1:])
			break
		}
		default: {
			fmt.Println("Unable to process request")
			response = "Invalid request. Available commands: 'squares', 'scramble', 'rsum', 'mean', 'validpass'"
			break
		}
	}

	response = strings.TrimSpace(response) + "\n"

	return response
}
