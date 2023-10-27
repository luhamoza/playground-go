package generics

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (c *CustomMap[K, V]) Insert(k K, v V) error {
	c.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func Foo[T any, Y any](val T, num Y) {
	fmt.Println(val, num)
}

func Generics() {
	m1 := NewCustomMap[int, string]()
	_ = m1.Insert(1, "Foo")
	_ = m1.Insert(2, "Bar")

	m2 := NewCustomMap[float64, int]()
	_ = m2.Insert(5.89, 3)
	_ = m2.Insert(3.97, 3)

	Foo[string, int]("Hello from Generics", 232)
}
