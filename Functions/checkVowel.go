package function

func CheckVowel(line []string) []string {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == "a" {
			for j := i + 1; j < len(line); j++ {
				if line[j] != "" && IsVowel(rune(line[j][0])) {
					line[i] = "an"
					break
				}
			}
		} else if line[i] == "A" {
			for j := i + 1; j < len(line); j++ {
				if line[j] != "" && IsVowel(rune(line[j][0])) {
					line[i] = "An"
					break
				}
			}
		} else if line[i] == "an" {
			for j := i + 1; j < len(line); j++ {
				if line[j] != "" && !IsVowel(rune(line[j][0])) {
					line[i] = "a"
					break
				}
			}
		} else if line[i] == "An" {
			for j := i + 1; j < len(line); j++ {
				if line[j] != "" && !IsVowel(rune(line[j][0])) {
					line[i] = "A"
					break
				}
			}
		}
	}
	return line
}
