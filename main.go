package main

import (
	"fmt"
	"playground/enums"
	"playground/functions"
	"playground/generics"
	"playground/interfaces"
	"playground/play"
	"playground/pointers"
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
}
