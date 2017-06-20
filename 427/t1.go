package main

import (
	"fmt"
)

func main() {
	str := "Hello,时间"
	for i, ch := range str{
		fmt.Println(i, ch)
	}
	var a, b = 1, 2
	fmt.Println(a, b)
}