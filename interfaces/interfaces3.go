package interfaces

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	name string
}

type Manager struct {
	name string
}

func (e *Engineer) GetName() string {
	return e.name
}

func (m *Manager) GetName() string {
	return m.name
}

func PrintDetails(e Employee) {
	fmt.Printf("Name: %s\n", e.GetName())
}

func Interfaces3() {
	e := &Engineer{
		name: "John",
	}
	m := &Manager{
		name: "Hilary",
	}
	PrintDetails(e)
	PrintDetails(m)
}
