package main

import "time"

func main() {
	done := make(chan bool, 1)

	go func() {
		println("goroutine exit.")
		time.Sleep(time.Second * 3)
		done <- true
	}()
	<-done
	println("main exit.")
}
