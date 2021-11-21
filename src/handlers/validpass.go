package handlers

import (
	"strings"
	"unicode"
)

func isValid(token string) bool {
	numChars := 0
	var upper, lower, number, special, minLength bool

	for _, char := range token {
		numChars++
		switch {
			case unicode.IsLower(char):
				lower = true
			case unicode.IsUpper(char):
				upper = true
			case unicode.IsNumber(char):
				number = true
			case unicode.IsSymbol(char) || unicode.IsPunct(char):
				special = true
		}
	}
	minLength = numChars >= 10

	return upper && lower && number && special && minLength
}

func ValidPassword(tokens []string) string {

	var validPasswords []string

	if len(tokens) < 1 {
		return "Provide at least one arg"
	} else {
		for _, token := range tokens {
			if isValid(token) {
				validPasswords = append(validPasswords, token)
			}
		}
	}

	if len(validPasswords) == 0 {
		return "No valid passwords found (a valid password should have a min length of 10 chars and contain: lowercase chars, uppercase chars, digits and symbols)"
	}

	return "Valid passwords: " + strings.Join(validPasswords, " ")
}
