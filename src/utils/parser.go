package utils

import "strconv"

type ParsedNum struct {
	Value int
	Token string
}

func ParseNum(token string) *ParsedNum {

	var num string

	for _, char := range token {
		if char >= '0' && char <= '9' {
			num += string(char)
		}
	}

	if len(num) > 0 {
		result, _ := strconv.Atoi(num)
		return &ParsedNum{Value: result, Token: token}
	}

	return nil
}
