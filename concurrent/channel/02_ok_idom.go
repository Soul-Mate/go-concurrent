package main

import (
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(2)
	go func() {
		send(c)
		defer close(c)
		wg.Done()
	}()
	go func() {
		recv(c)
		wg.Done()
	}()
	wg.Wait()
}

func send(c chan <- int)  {
	for i := 0; i < 5;i++  {
		time.Sleep(time.Second)
		c<-i
	}
}

func recv(c <-chan int) {
	for {
		x, ok := <-c
		if !ok {
			println(x, ok)
			break;
		}
		println(x, ok)
	}
}