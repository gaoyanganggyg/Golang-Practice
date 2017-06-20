package main

import (
	"time"
	"fmt"
)

func show(delay time.Duration)  {
	for {
		for _, x := range `-+-+` {
			fmt.Printf("\r%c", x)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(time.Now().Format("15:04:05\n"))
	//go show(100*time.Millisecond)
	//const N  = 45
	//fibN := fib(N)
	//fmt.Printf("\rFibonacci(%d) = %d\n", N, fibN)
}
