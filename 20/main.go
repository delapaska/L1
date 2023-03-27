package main

import (
	"fmt"
	"strings"
)

func ReverseStrWords(str string) {
	newStr := strings.Split(str, " ")
	j := len(newStr) - 1

	for i := 0; i < len(newStr)/2; i++ {
		newStr[i], newStr[j] = newStr[j], newStr[i]
		j--
	}

	fmt.Println(newStr)
}

func main() {

	ReverseStrWords("sun dog snow")
}
