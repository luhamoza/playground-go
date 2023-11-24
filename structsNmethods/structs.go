package structsNmethods

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

func (p *Person) Print() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Job: %s\n", p.job)
}

func (p *Person) ModifyAge(age int) int {
	age += 10
	return age
}

func Structs() {
	person := &Person{
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
	person.Print()
	person.age = person.ModifyAge(person.age)
	fmt.Println("New Age")
	person.Print()
}
