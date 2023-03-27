package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}
	// Объявляем набор горутин
	var wg sync.WaitGroup
	wg.Add(len(arr))

	// Через горутины

	for _, elements := range arr {
		// создаём горутины
		go numberSquares(elements, &wg)
	}
	wg.Wait()

	// Через каналы
	/*
		chnl := make(chan int, len(arr))

		defer close(chnl)

		for _, elem := range arr {
			go squareWithChannel(elem, chnl)
		}

		for range arr {
			fmt.Println(<-chnl)
		}
	*/

}

func numberSquares(number int, wg *sync.WaitGroup) {
	// Вывод квадрата числа
	fmt.Println(number * number)
	// Уменьшение количества горутин
	wg.Done()
}

func squareWithChannel(number int, chanelInt chan int) {
	//Передаем в канал квадрат числа
	chanelInt <- number * number
}
