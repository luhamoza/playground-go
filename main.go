package main

import (
	"fmt"
	"playground/enums"
	"playground/functions"
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

	fmt.Println("Functions:")
	functions.Functions()
	fmt.Println()

	fmt.Println("Interfaces2:")
	interfaces.Interface2()
	fmt.Println()
}
