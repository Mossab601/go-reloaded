package function

import (
	"strings"
	"unicode"
)

func Up(str string) string {
	res := ""
	for _, c := range str {
		res += string(unicode.ToUpper(c))
	}
	return res
}

func Low(str string) string {
	res := ""
	for _, c := range str {
		res += string(unicode.ToLower(c))
	}
	return res
}

func Cap(str string) string {
	res := ""
	for i, c := range str {
		if i == 0 {
			res += string(unicode.ToUpper(c))
		} else {
			res += string(unicode.ToLower(c))
		}
	}
	return res
}

func ArrayToString(arr [][]string) string {
	var result string
	for i, row := range arr {
		result += strings.Join(row, " ")
		if i < len(arr)-1 {
			result += "\n"
		}
	}

	return result
}