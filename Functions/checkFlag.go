package function

import (
	"strconv"
	"strings"
)

func Flag(line []string) []string {
	result := []string{}
	for i := 0; i < len(line); i++ {
		word := line[i]
		word == "(up)" || word == "(low)" || word == "(cap)" || word == "(bin)" || word == "(hex)" {
			continue
		}
		for i < len(line)-1 {
			next := line[i+1]
			if next == "(up)" || next == "(low)" || next == "(cap)" || next == "(bin)" || next == "(hex)" {
				word = ApplyFlag(word, next)
				i++
			} else {
				break
			}
		}
		result = append(result, word)
	}
	return result
}

func FlagWithParams(line []string) []string {
	for i := 0; i < len(line)-1; i++ {
		_, err1 := strconv.Atoi(line[i+1][:len(line[i+1])-1])
		if strings.HasPrefix(line[i], "(") && strings.HasSuffix(line[i], ",") && err1 == nil {
			line[i] = line[i][:len(line[i])-1] + ")"
		}

		if strings.HasSuffix(line[i+1], ")") {
			if line[i] == "(up)" || line[i] == "(low)" || line[i] == "(cap)" {
				if num, err := strconv.Atoi(line[i+1][:len(line[i+1])-1]); err == nil {
					remainingWords := line[i+2:]
					line = ApplyFlagWithNumber(line[:i], line[i], num)
					line = append(line, remainingWords...)
				}
			}
		}
	}
	return line
}
