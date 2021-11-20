package handlers

func Scramble(tokens []string) string {
	var result string

	if len(tokens) < 2 {
		return "Provide at least 2 args"
	} else {
		length := len(tokens[0])
		for i := 0; i < length; i++ {
			var scrambledToken string
			for _, token := range tokens {
				if length != len(token){
					return "Provided args must have equal lengths"
				}
				scrambledToken += string(token[i])
			}
			result += scrambledToken + " "
		}
	}

	return result
}

