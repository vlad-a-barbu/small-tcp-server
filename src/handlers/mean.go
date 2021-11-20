package handlers

import (
	"fmt"
	"strconv"
)

func sumDigits(num int) int {
	remainder := 0
	sum := 0

	for num != 0 {
		remainder = num % 10
		sum += remainder
		num = num / 10
	}
	return sum
}
func ArithmeticMean(tokens []string) string {
	var result string

	if len(tokens) < 3 {
		return "Provide at least three args"
	} else {

		var sumNums int
		var cnt int

		lowerLimit, err := strconv.Atoi(tokens[0])
		if err != nil {
			return "Every arg should be an integer"
		}

		upperLimit, err := strconv.Atoi(tokens[1])
		if err != nil {
			return "Every arg should be an integer"
		}

		for _, token := range tokens[2:] {
			num, err := strconv.Atoi(token)
			if err != nil {
				return "Every arg should be an integer"
			}

			sumDig := sumDigits(num)
			if sumDig >= lowerLimit && sumDig <= upperLimit {
				sumNums += num
				cnt++
			}
		}

		mean := sumNums / cnt

		return fmt.Sprintf("Mean = %d", mean)
	}

	return result
}
