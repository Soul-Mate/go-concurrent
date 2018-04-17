package main

import "sync"

func main()  {
	s:=sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		s.Add(1)
		go func(n int) {
			println("goroutine: ", n)
			s.Done()
		}(i)
	}
	s.Wait()
}