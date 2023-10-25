package interfaces

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongodbNumberStorer struct {
	//something values here
}

type ApiServer struct {
	numberStore NumberStorer
}

func (m MongodbNumberStorer) GetAll() ([]int, error) {
	return []int{3, 5, 6}, nil
}

func (m MongodbNumberStorer) Put(number int) error {
	fmt.Println("Number store in MongoDB")
	return nil
}

func interfaces() {

	apiServer := ApiServer{numberStore: MongodbNumberStorer{}}
	exampleN, err := apiServer.numberStore.GetAll()
	if err != nil {
		fmt.Println("These some error")
		return
	}
	fmt.Println(exampleN)
}
