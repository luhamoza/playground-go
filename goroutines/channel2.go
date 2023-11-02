package goroutines

import (
	"fmt"
	"time"
)

func fetchResourceThree(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result: %d", n)
}
func Channel2() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	msgch <- "D"

	//msg := <-msgch
	//fmt.Println("The msg is:", msg)
	//msg = <-msgch
	//fmt.Println("The msg is:", msg)

	//refactor code above
	close(msgch)
	for msg := range msgch {
		fmt.Println("The msg is:", msg)
	}

	fmt.Println("done reading all the messages from the channel")
}
