package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Вместе с Main
func waitMain() {
	go func() {
		for {
		}
	}()
}

//с помощью return

func ret() {
	go func() {

		return
	}()
}

// WaitGroup
func WG() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()
}

// Используя каналы
func do_stuff() int {
	return rand.Intn(100)
}

func channel() {
	ch := make(chan int, 100)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- do_stuff():
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("receive value: ", i)
	}

	fmt.Println("finish")
}

// Используя контекст
func ctx() {
	forever := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				forever <- struct{}{}
				return
			default:
				fmt.Println("for loop")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}

func main() {
	waitMain()

}
