package main

import (
	"runtime"
	"time"
)

func main()  {
	go func() {
		defer println("defer 1.")
		defer println("defer 2.")
		func(){
			defer func() {
				println("recover() == ", recover())
			}()
			println("call runtime.Goexit: ")
			runtime.Goexit()
			println("a")
		}()
		println("b")
	}()
	time.Sleep(time.Second * 2)
}
