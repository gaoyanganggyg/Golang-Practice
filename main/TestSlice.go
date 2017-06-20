package main

import "fmt"

func t1(sl []string)  {
	for i:=0; i<len(sl); i++{
		sl[i] = sl[i] + "sd"
	}

	fmt.Println(sl)
}

func t2(m map[string]int)  {
	for k, v := range m{
		m[k] = 2*v
	}
}

func t3(a *[2]int)  {
	for i:=0; i<len(a); i++{
		a[i] = 2*a[i]
	}
	fmt.Println(a)
}


func main() {
	//m := map[string]int{"a": 1, "b": 10}
	//fmt.Println(m)
	//t2(m)
	//fmt.Println(m)
	a := [2]int{1, 2}
	fmt.Println(a)
	t3(&a)
	fmt.Println(a)
}
