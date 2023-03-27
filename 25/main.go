package main

import (
	"fmt"
	"time"
)

// Функция реализовывает Sleep в задержках, равным секундам.
// Для того, чтобы время было более гибким,
// можно заменить time.Second на time.Millisecond
func Sleep(x int) {

	<-time.After(time.Second * time.Duration(x))
}

func main() {

	var str string
	str = "Hello World!"

	var x int

	fmt.Println("Введите время для задержки между выводами в секундах")
	fmt.Scan(&x)
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
		Sleep(x)
	}

}
