package main

import (
	"fmt"
	"playground/enums"
	"playground/functions"
	"playground/generics"
	"playground/goroutines"
	"playground/interfaces"
	"playground/play"
	"playground/pointers"
	"playground/test"
	"playground/types"
)

func main() {
	fmt.Println("Functions:")
	play.Play()
	play.RunTime()
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

	fmt.Println("Generics:")
	generics.Generics()
	fmt.Println()

	fmt.Println("Pointers:")
	pointers.Pointer()
	fmt.Println("Pointers2:")
	fmt.Println()

	fmt.Println("Pointers3:")
	pointers.Pointer3()
	fmt.Println()

	fmt.Println("Test:")
	test.Example()
	fmt.Println()

	fmt.Println("Goroutines:")
	goroutines.Goroutines()
	fmt.Println()

	fmt.Println("Channel:")
	goroutines.Channel()
	fmt.Println()

	fmt.Println("Channel 2:")
	goroutines.Channel2()
	fmt.Println()

}
