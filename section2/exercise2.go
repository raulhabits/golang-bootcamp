package section2

import "fmt"

type Musician struct {
	Name string
}

type MusicalPlayer interface {
	Greetings() string
}

type Trumpeter struct {
	Musician
}

type Violinist struct {
	Musician
}

func (receiver *Violinist) Greetings() string {
	return fmt.Sprintf("Hello %s are you going to do a violin solo today?\n", receiver.Name)
}
func (receiver *Trumpeter) Greetings() string {
	return fmt.Sprintf("Hi %s this is a good day for practicing with your trumpet?\n", receiver.Name)
}



