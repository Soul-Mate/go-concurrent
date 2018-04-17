package main

import "time"

func main()  {
	go func() {
		println("hello goroutine!")
	}()
	time.Sleep(time.Second)
}
