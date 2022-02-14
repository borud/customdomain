package customdomain

import "fmt"

type Sometype struct{}

func (s *Sometype) Hello() {
	fmt.Println("hello there")
}
