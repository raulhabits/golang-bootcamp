package section1

import (
	"fmt"
	"testing"
)

func TestSlice1_addIndex(t *testing.T) {
	slice1 := Slice1{
		data: []int{1, 2 , 3, 4},
	}

	fmt.Println("Slice before being call addIndex function Slice1 ", slice1.data)

	slice1.addIndex(5)

	fmt.Println("Slice after call addIndex function Slice1 ", slice1.data)

	slice2 := Slice2{
		data: []int{1, 2 , 3, 4},
	}

	fmt.Println("Slice before being call addIndex function Slice2 ", slice2.data)

	slice2.addIndex(5)

	fmt.Println("Slice after call addIndex function Slice2 ", slice2.data)

}
