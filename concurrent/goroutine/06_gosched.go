package main

import (
	"runtime"
	"sync"
	"time"
)

func main()  {
	s := sync.WaitGroup{}
	s.Add(1)
	go func() {
		wg := sync.WaitGroup{}
		go func() {
			println("b")
			time.Sleep(time.Second * 3)
			wg.Done()
		}()
		wg.Add(1)
		for i := 0; i < 5; i++ {
			if i == 1 {
				runtime.Gosched()
			}
			println("a: ", i)
		}
		wg.Wait()
		s.Done()
	}()
	s.Wait()
}