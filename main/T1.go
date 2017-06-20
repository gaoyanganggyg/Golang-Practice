package main

import (
	"fmt"
	"time"
)


//
//import "fmt"
//
//type Integer int
//
//func (a *Integer) Less(b Integer) {
//	*a += b
//}
//
//func main() {
//	var a Integer = 10
//	a.Less(20)
//	fmt.Println(a)
//
//}

//var count int = 0
//
//func Count(lock *sync.Mutex)  {
//	lock.Lock()
//	count++
//	lock.Unlock()
//}
//
//func main()  {
//	lock := &sync.Mutex{}
//	for i:=0; i<10; i++ {
//		go Count(lock)
//	}
//	for  {
//		lock.Lock()
//		c := count
//		lock.Unlock()
//		runtime.Gosched()
//		if c >= 10 {
//			break
//		}
//	}
//}

func main() {
	tick := time.Tick(2000 * time.Millisecond)
	//for {
		select {
		case <-tick:
			fmt.Println("asd")

		}
	//}
}