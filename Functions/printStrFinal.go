package function

func PrintStrFinal(result string) string {
	res := Split(result)
	arr := ArrayToString(res)
	ignPonct := CheckPonct(arr)
	singlQot := CheckQuotes(ignPonct)
	return singlQot
}
