package main

import (
	"fmt"
	"time"
)

// goroutine is processed in the same address space,
// which means that must exclusive shared address.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
