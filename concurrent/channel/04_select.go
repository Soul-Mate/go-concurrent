package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c1 := make(chan int)
	c2 := make(chan int)
	wg.Add(2)
	go func() {
		for {
			select {
			case a := <-c1:
				println("a: ", a)
			case b := <-c2:
				println("b: ", b)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second)
			c2 <- i * 10
		}
		close(c1)
		close(c2)
		wg.Done()
	}()
	wg.Wait()
}
