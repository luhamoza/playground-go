package structs

import "fmt"

type Person struct {
	name      string
	age       int
	job       string
	isWorking bool
	Project   Project
}

type Project struct {
	isValid bool
	isAdmin bool
	tasks   int
	tech    []string
}

func Structs() {
	person := Person{
		name:      "John",
		age:       23,
		job:       "Engineer",
		isWorking: false,
		Project: Project{
			isValid: false,
			isAdmin: false,
			tasks:   0,
			tech:    []string{"Go, PostgresSQL, CSS"},
		},
	}
	fmt.Printf("%+v\n", person)
}
