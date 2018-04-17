package main

import "sync"

type dataType int

var v dataType

var goTLS = []struct{
	id int
	data dataType
}{}

func main()  {
	s := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		s.Add(1)
		go func(n int, cv dataType) {
			for j := 0; j < 100; j++ {
				cv++
			}
			goTLS = append(goTLS, struct {
				id   int
				data dataType
			}{id: n, data: cv})
			println("goroutine ", n, ": done.")
			s.Done()
		}(i, v)
	}
	s.Wait()
	merge()
	println("merge v = ", v)
}

func merge() {
	for _, s := range goTLS {
		println("merge ", s.id, " goroutine data.")
		v += s.data
	}
}