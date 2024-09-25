package function

func CheckPonct(s string) string {
	var output []rune
	var prevRune rune
	myStr := s
	for index, r := range myStr {

		if IsPonc(r) {
			if prevRune == ' ' {
				output = output[:len(output)-1]
			}
			output = append(output, r)
			if index != len(myStr)-1 {
				if !IsPonc(rune(myStr[index+1])) && rune(myStr[index+1]) != ' ' {
					output = append(output, ' ')
				}
			}

		} else {
			output = append(output, r)
		}

		prevRune = r
	}

	return string(output)
}
