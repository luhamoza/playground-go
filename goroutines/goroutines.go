package goroutines

import (
	"fmt"
	"time"
)

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
func Goroutines() {
	result := fetchResource(2)
	fmt.Println(result)
	//change above code to goroutines
	//syntax 1
	go fetchResource(2)

	//syntax 2
	go func() {
		fetchResource(2)
	}()

	// channel
	// 1 unbuffered channel
	// 2 buffered channel
	//resultch := make(chan string) // <- unbuffered channel
	//resultch <- "foo"             // <- is now FULL -> it will block -> block here

	//this code will never execute
	//result := <-resultch

	//fmt.Println(result)

	// channel.go solution to make the code work
}
