package handlers

import (
	"fmt"
	"strconv"
)

func reverse(num int) int {
	var result int
	for num > 0 {
		remainder := num % 10
		result = (result * 10) + remainder
		num /= 10
	}
	return result
}

func ReversedSum(tokens []string) string {

	var result string

	if len(tokens) < 1 {
		return "Provide at least one arg"
	} else {
		sum := 0
		for _, token := range tokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				return "Every arg should be an integer"
			}
			sum += reverse(num)
		}
		result = fmt.Sprintf("Reversed elements sum = %d", sum)
	}

	return result
}
