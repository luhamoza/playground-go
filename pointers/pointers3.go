package pointers

import "fmt"

type Database struct {
	user string
}
type Server struct {
	db *Database
}

func (s *Server) GetUser() string {
	if s.db == nil {
		return "DB is not initialize"
	}
	return s.db.user
}

func Pointer3() {
	s := &Server{}
	fmt.Println(s.GetUser())
}
