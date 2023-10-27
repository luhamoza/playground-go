package interfaces

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongodbNumberStorer struct {
	// something values here
}
type PostgresNumberStorer struct {
	// some value
}

type ApiServer struct {
	numberStore NumberStorer
}

func (m MongodbNumberStorer) GetAll() ([]int, error) {
	return []int{3, 5, 6}, nil
}

func (m MongodbNumberStorer) Put(number int) error {
	fmt.Println("Number stored in MongoDB")
	return nil
}
func (p PostgresNumberStorer) GetAll() ([]int, error) {
	return []int{8, 9, 6, 2, 5, 5, 7}, nil
}

func (p PostgresNumberStorer) Put(number int) error {
	fmt.Println("Number stored in Postgres")
	return nil
}

func Interfaces() {
	
	//mongodb
	apiServer := ApiServer{numberStore: MongodbNumberStorer{}}
	example, err := apiServer.numberStore.GetAll()
	if err != nil {
		fmt.Println("These some error")
		return
	}
	fmt.Println(example)

	//postgres
	apiServerTwo := ApiServer{numberStore: PostgresNumberStorer{}}
	var (
		errTwo     error
		exampleTwo []int
	)
	if exampleTwo, errTwo = apiServerTwo.numberStore.GetAll(); errTwo != nil {
		fmt.Println("There some error")
		return
	}
	fmt.Println(exampleTwo)
}
