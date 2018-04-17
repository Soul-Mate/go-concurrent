package main

import "runtime"

func main() {
	cpu := runtime.NumCPU()
	println("current system cpu: ", cpu)
	cpu = runtime.GOMAXPROCS(cpu)
	println("set runtime.GOMAXPROCS cpu: ", cpu)
}
