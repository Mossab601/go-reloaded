package function

import (
	"strconv"
)

func ConvDic(str, base string) int64 {
	if base == "(hex)" {
		res, _ := strconv.ParseInt(str, 16, 64)
		return res
	} else if base == "(bin)" {
		res, _ := strconv.ParseInt(str, 2, 64)
		return res
	}
	return 0
}

func IsHexBin(str, base string) (bool, string) {
	var baseToConverte []rune
	if base == "(bin)" {
		baseToConverte = []rune{'-', '+', '0', '1'}
	} else if base == "(hex)" {
		baseToConverte = []rune{'-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F'}
	}

	for _, c := range str {
		IsValid := false
		for _, validChar := range baseToConverte {
			if c == validChar {
				IsValid = true
				break
			}
		}
		if !IsValid {
			return false, ""
		}
	}
	return true, str
}

func ConvertOrReturnString(str, base string) string {
	isValid, validStr := IsHexBin(str, base)
	if !isValid {
		return str
	}
	convertedValue := ConvDic(validStr, base)
	return strconv.FormatInt(convertedValue, 10)
}
