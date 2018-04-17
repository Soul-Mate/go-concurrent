package main

import "sync"

func main()  {
	c :=  make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()

	for x := range c {
		println("x = ", x)
	}
}