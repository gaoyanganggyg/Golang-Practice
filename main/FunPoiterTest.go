package main

import (
	"fmt"
)

type A int

func (a *A) T1(b A)  (c A){
	c = *a + b
	return
}

func main() {
	a := A(12)
	fmt.Println(a.T1(A(10)))


}
