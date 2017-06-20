package main

import (
	"fmt"
)

const (
	S = 1
	B = 2
)


func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(S)
	fmt.Println(B)
	var a *int
	fmt.Println(a == nil)
}