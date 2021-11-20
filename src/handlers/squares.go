package handlers

import (
	"../utils"
	"fmt"
	"math"
	"strconv"
)

func isSquare(num int) bool {
	return math.Sqrt(float64(num)) == math.Floor(math.Sqrt(float64(num)))
}

func getSquares(nums []utils.ParsedNum) []utils.ParsedNum {
	var squares []utils.ParsedNum

	for _, num := range nums {
		if isSquare(num.Value) {
			squares = append(squares, num)
		}
	}

	return squares
}

func Squares(tokens []string) string {
	var result string

	if len(tokens) < 1 {
		return "Provide at least one arg"
	} else {

		var nums []utils.ParsedNum

		for _, token := range tokens {
			num := utils.ParseNum(token)
			if num != nil {
				nums = append(nums, *num)
			}
		}

		squares := getSquares(nums)
		numSquares := len(squares)

		if numSquares > 0 {
			result = strconv.Itoa(numSquares) + " square nums found: "
			for _, square := range squares {
				result += fmt.Sprintf("%d from %s; ", square.Value, square.Token)
			}
		} else {
			return "No squares were found"
		}
	}

	return result
}
