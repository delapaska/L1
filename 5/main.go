package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Отправляет в канал данные с интервалом в 1 секунду
// Данные типа Int в интервале от 1 до 1000
func Sender(ch chan int) {
	for {
		ch <- rand.Intn(1000)
		//time.Sleep(time.Second * 1)
	}
}

// Считывание и вывод данных
func Reader(ch chan int) {

	for {
		tmp := <-ch
		fmt.Println(tmp)
	}
}

func main() {
	ch := make(chan int)
	var N int
	fmt.Println("Введите число N:")
	fmt.Scanf("%d", &N)
	go Sender(ch)
	go Reader(ch)
	time.Sleep(time.Second * (time.Duration(N)))
	close(ch)

}
