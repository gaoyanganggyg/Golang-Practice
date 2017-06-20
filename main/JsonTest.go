package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	//var d  = &Person{"Gao", 18, "M"}
	d1 := map[string]interface{}{"name": "XiaoMing", "age": 10}

	bt1, _ := json.Marshal(d1)
	fmt.Println(string(bt1))
	var res1 interface{}
	err := json.Unmarshal(bt1, &res1)
	if nil != err {
		fmt.Println(err)
		return
	}
	if res2, ok := res1.(map[string]interface{}); ok {
		fmt.Println(res2["name"])
		fmt.Println(res2["age"])
		fmt.Println(res2["data"])
	}
}
