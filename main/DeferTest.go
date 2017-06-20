package main

import (
	"time"
	"log"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the
	// ...lots of work…
	time.Sleep(4 * time.Second) // simulate slow
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg,time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
