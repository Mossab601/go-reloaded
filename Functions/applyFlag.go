package function

func ApplyFlag(word, flag string) string {
	switch flag {
	case "(up)":
		return Up(word)
	case "(low)":
		return Low(word)
	case "(cap)":
		return Cap(word)
	case "(hex)":
		return ConvertOrReturnString(word, "(hex)")
	case "(bin)":
		return ConvertOrReturnString(word, "(bin)")
	default:
		return word
	}
}

func ApplyFlagWithNumber(words []string, flag string, count int) []string {
	for i := len(words) - 1; i >= 0 && count > 0; i-- {
		words[i] = ApplyFlag(words[i], flag)
		count--
	}
	return words
}
