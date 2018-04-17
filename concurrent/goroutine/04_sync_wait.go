package main

import (
	"sync"
	"time"
)

func main() {
	done := make(chan bool,1)
	s := sync.WaitGroup{}
	s.Add(1)
	go func() {
		println("i'am start. (1)")
		time.Sleep(time.Second * 2)
		println("i'am done. (1)")
		s.Done()
	}()

	go func() {
		s.Wait()
		println("i'am start. (2)")
		time.Sleep(time.Second * 2)
		println("i'am done. (2)")
		done <- true
	}()
	<-done
}