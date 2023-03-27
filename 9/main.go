package main

import "fmt"

func readInt(x *int) {
	var y int = 0
	fmt.Scan(&y)
	*x = y
}
func main() {

	// Ввод размера массива и его содержимого
	var n int = 0
	fmt.Println("Введите количество элементов:")
	readInt(&n)
	arr := make([]int, n)
	fmt.Println("Введите элементы:")
	var k int = 0
	for i := 0; i < n; i++ {
		readInt(&k)
		arr[i] = k
	}

	// Работа с каналами
	var fChannel = make(chan int)
	var sChannel = make(chan int)

	go func() {
		for _, element := range arr {
			fChannel <- element
		}
	}()
	go channelWork(fChannel, sChannel)
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-sChannel)
	}
}

func channelWork(f chan int, s chan int) {
	for {
		x, ok := <-f
		if !ok {
			break
		}
		s <- x * 2

	}
}
