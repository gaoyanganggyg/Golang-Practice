package main

import (
	"fmt"
)

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}
// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

//相邻两位反转
var tmp *IntList = nil
var head *IntList = nil
func Swap(l *IntList) *IntList{
	for nil != l && nil != l.Tail{
		i1 := l.Tail
		t := i1.Tail
		i1.Tail = l
		l.Tail = t
		if nil == head{
			head = i1
		}
		if nil != tmp{
			tmp.Tail = i1
		}
		tmp = l
		l = l.Tail
	}
	return head
}

//链表反转
func Reverse(l *IntList) *IntList{
	var before *IntList = nil
	for l != nil{
		t := l.Tail
		l.Tail = before
		before = l
		l = t
	}
	return before
}


func main() {
	l := &IntList{Value:1}
	head := l
	for x:=2; x<=10; x++{
		l.Tail = &IntList{Value:x}
		l = l.Tail
	}

	tmp:=head
	for tmp!=nil{
		fmt.Println(tmp.Value)
		tmp = tmp.Tail
	}

	fmt.Println("--------------------------")
	tmp = Reverse(head)
	//tmp = Swap(head)
	for tmp!=nil{
		fmt.Println(tmp.Value)
		tmp = tmp.Tail
	}


	//fmt.Println(head.Sum())
}
