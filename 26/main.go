package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем map для хранения количества символов
	chars := make(map[rune]int)

	// Проходимся по строке и добавляем символы в map
	for _, c := range str {
		chars[c]++
		if chars[c] > 1 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(isUnique(s1)) // true
	fmt.Println(isUnique(s2)) // false
	fmt.Println(isUnique(s3)) // false
}
