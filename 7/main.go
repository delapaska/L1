package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var m = make(map[int]int)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			m[i] = i * 2
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	for k, v := range m {
		fmt.Printf("%d:%d\n", k, v)
	}
}
