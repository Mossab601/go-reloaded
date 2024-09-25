package function

func CheckQuotes(text string) string {
	result := ""
	first := false
	next := false
	for i, char := range text {
		if char == '\'' {
			if i == 0 {
				result += string(char)
				first = true
				if i+1 < len(text) && text[i+1] == ' ' {
					next = true
				}
			} else if i == len(text)-1 && !first {
				result += "'"
			} else {

				if i > 0 && i < len(text)-1 &&
					IsAlphaNumeric(rune(text[i-1])) && IsAlphaNumeric(rune(text[i+1])) {
					result += "'"
					continue
				}
				if !first {
					if result[len(result)-1] != ' ' {
						result += " "
					}
					result += "'"
					if text[i+1] == ' ' {
						next = true
					}
					first = true
				} else {
					first = false
					if result[len(result)-1] == ' ' {
						result = result[:len(result)-1]
					}
					result += "'"

				}
			}
		} else {
			if next {
				next = false

				continue
			}
			result += string(char)
		}

	}
	return result
}
