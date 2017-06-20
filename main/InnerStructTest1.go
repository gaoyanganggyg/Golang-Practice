package main

import (
	"sync"
	"fmt"
	"strings"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping:make(map[string]string)}


func Lookup(key string) (res string) {
	cache.Lock()
	res = cache.mapping[key]
	cache.Unlock()
	if "" == strings.TrimSpace(res){
		res = "nil"
	}
	return
}

func main() {
	fmt.Println(Lookup("a"))
}