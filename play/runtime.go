package play

import (
	"fmt"
	"runtime"
)

func RunTime() {

	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux Machine")
	case "darwin":
		fmt.Println("MacOs")
	default:
		fmt.Println("windows")
	}

}
