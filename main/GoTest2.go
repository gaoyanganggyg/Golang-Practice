package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i:=0; i<10; i++{
			c1<-i
		}
		close(c1)
	}()

	go func() {
		for x := range c1{
			c2 <- x*x
		}
		close(c2)
	}()

	for x := range c2{
		fmt.Println(x)
	}
}