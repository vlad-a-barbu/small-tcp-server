package utils

import "strings"

func Tokenize(message string) []string {

	tokens := strings.Split(strings.TrimSpace(message), " ")

	var validTokens []string
	for _, token := range tokens {
		if len(token) > 0 {
			validTokens = append(validTokens, token)
		}
	}

	return validTokens
}
