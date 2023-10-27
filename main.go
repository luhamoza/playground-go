package main

import (
	"fmt"
	"playground/enums"
	"playground/interfaces"
	"playground/play"
	"playground/types"
)

func main() {
	fmt.Println("Functions:")
	play.Play()
	fmt.Println()

	fmt.Println("Enums:")
	enums.Enums()
	fmt.Println()

	fmt.Println("Interfaces:")
	interfaces.Interfaces()
	fmt.Println()

	fmt.Println("Types:")
	types.Types()
	fmt.Println()
}
