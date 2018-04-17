package main

func main() {
	println("goroutine returned: ", test())
}

func test() int {
	var ret int
	done := make(chan int, 1)
	go func() {
		for i := 1; i <= 100; i++ {
			ret += i * 10
		}
		done <- ret
	}()
	return <-done
}
