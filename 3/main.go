package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}
	/*
		// Через Горутины
		sum := 0
		//объявляем набор горутин
		var wg sync.WaitGroup
		wg.Add(len(arr))

		for _, elements := range arr {
			elements := elements
			go func() {
				//вычисляем сумму
				sum += (elements * elements)
				//уменьшаем счётчик горутин
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println(sum)
	*/

	// С помощью mutex

	var mu sync.Mutex
	sum := 0
	var num int
	for _, el := range arr {
		num = el
		go func() {
			mu.Lock()
			sum += (num * num)
			mu.Unlock()
		}()
	}
	fmt.Println(sum)

}
