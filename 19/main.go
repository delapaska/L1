package main

import "fmt"

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}

	return

}
func ReverseRune(str string) {
	newRune := []rune(str)
	j := len(newRune) - 1

	for i := 0; i < len(newRune)/2; i++ {
		newRune[i], newRune[j] = newRune[j], newRune[i]
		j--
	}

	fmt.Println(string(newRune))
}

func main() {
	fmt.Println(Reverse("главрыба"))
	ReverseRune("главрыба")
}
