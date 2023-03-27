package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count++
}

func main() {

	obj := Counter{0}
	var mutex sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(4)

	//i должно быть равно числу горутин
	for i := 0; i < 4; i++ {
		go func() {
			// любое значение i
			for i := 0; i < 10; i++ {
				mutex.Lock()
				obj.Inc()
				mutex.Unlock()
			}
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println(obj.count)
}
