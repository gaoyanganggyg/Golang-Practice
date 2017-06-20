package main

import "fmt"

func Ptest(v ...int)  {
	var res  = 0
	for _, x := range v{
		res += x
	}
	fmt.Println(res)
}

func main() {
	Ptest(1, 2, 3)
	Ptest(1, 2, 3, 90, 45)
}
