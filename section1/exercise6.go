package section1

import "fmt"

type Slice1 struct {
	data []int
}

type Slice2 struct {
	data []int
}

func (s Slice1) addIndex(value int)  {
	s.data = append(s.data, value)
	fmt.Println("Slice on the addIndex function on Slice1 ", s.data)
}


func (s *Slice2) addIndex(value int)  {
	s.data = append(s.data, value)
	fmt.Println("Slice on the addIndex function on Slice2 ", s.data)
}