package main

import (
	"fmt"
)

func setBit(num int64, pos uint, val bool) int64 {
	if val {
		// Сдвиг влево на количество pos
		// Оператор или предназначен для сохранение значения если он был == 1
		// Иначе устанавливается 0
		num |= (1 << pos)
	} else {
		// установка i-го бита в 0
		num &= ^(1 << pos)
	}
	return num
}

func main() {
	var num int64 = 10 // число 1010 в двоичной системе счисления
	fmt.Printf("Число до изменений: %d\n", num)

	// устанавливаем 2-й бит в 1
	num = setBit(num, 3, true)

	fmt.Printf("Число после установки 2-го бита в 1: %d\n", num)

	// устанавливаем 1-й бит в 0
	num = setBit(num, 2, false)

	fmt.Printf("Число после установки 1-го бита в 0: %d\n", num)
}
