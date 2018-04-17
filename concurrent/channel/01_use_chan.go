package main

import "sync"

func main()  {
	done := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		done <- 1
		done <- 2
		wg.Done()
	}()
	go func() {
		println(<-done)
		wg.Done()
	}()
	wg.Wait()
}
