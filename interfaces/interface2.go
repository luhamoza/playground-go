package interfaces

import "fmt"

type Putter interface {
	Put(id int, val any) error
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}
type Server struct {
	store Storage
}
type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

func updateValue(id int, value any, p Putter) error {
	return p.Put(id, value)
}

func Interface2() {
	s := &Server{
		store: &FooStorage{}}
	v := &Server{
		store: &BarStorage{}}
	//a, _ := s.store.Get(2)
	fmt.Println(updateValue(1, "any", s.store))
	fmt.Println(updateValue(2, "bar", v.store))

}
