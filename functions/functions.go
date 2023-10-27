package functions

import (
	"fmt"
	"strings"
)

type FuncString func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}
func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Indexer(s string) FuncString {
	return func(v string) string {
		return s + v
	}
}
func TransformString(s string, fn FuncString) string {
	return fn(s)
}

func Functions() {
	fmt.Println(TransformString("Hello world", Uppercase))
	fmt.Println(TransformString("Hello world", Lowercase))
	fmt.Println(TransformString("Hello world", Indexer("Hey ")))
}
