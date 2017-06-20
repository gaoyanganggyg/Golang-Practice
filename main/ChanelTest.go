package main

import (
	"time"
	"fmt"
)

func T1()  {
	abort := make(chan struct{})

	go func() {
		time.Sleep(4*time.Second)
		abort<- struct{}{}
	}()


	select {
	case <-time.After(5*time.Second):
		fmt.Println("launch")
	case <-abort:
		fmt.Println("abort")
	}
	fmt.Println("Done")
}

func T2()  {
	c1 := make(chan int, 1)
	for i:=0; i<10; i++{
		select {
		case x:=<-c1:
			fmt.Println(x)
		case c1<-i:
		}
	}
}

func T4()  {
	abort := make(chan struct{})
	go func() {
		abort<-nil
	}()

	select {
	case <-abort:
		fmt.Println("Abort")
	//default:
	//	fmt.Println("Default")
	}
}


func main() {
	//T1()
	//T2()
	T4()
}