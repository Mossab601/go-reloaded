package main

import (
	"fmt"
	function "function/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("check numbers of arguments")
		return
	}
	if os.Args[1] != "sample.txt" || os.Args[2] != "result.txt" {
		fmt.Println("check your input & output")
		return
	}
	sample, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(sample)
	os.WriteFile(os.Args[2], []byte(strings.Trim(function.PrintStrFinal(result), " ")), 644)

}
