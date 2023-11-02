package goroutines

import (
	"fmt"
	"time"
)

func fetchResourceTwo(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Result: %d", n)
}

func Channel() {
	resultch2 := make(chan string)
	go func() {
		result2 := <-resultch2
		fmt.Println(result2)
	}()
	resultch2 <- "foo channel"
}
