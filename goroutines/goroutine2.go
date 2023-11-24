package goroutines

import (
	"fmt"
	"time"
)

func Hello(s string) {
	fmt.Println("Hello", s)
}

func GoRoutines2() {
	go Hello("Go Routines")
	fmt.Println("I will printed first")
	time.Sleep(time.Second * 1)
}
