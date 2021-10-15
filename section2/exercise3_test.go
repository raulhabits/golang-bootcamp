package section2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_performAreaAndPerimeter(t *testing.T) {
	shapes := []Shape{
		&Circle{2},
		&Rectangle{2, 3},
	}

	for _, shape := range shapes {
		area, perimeter := performAreaAndPerimeter(shape)
		fmt.Println(fmt.Sprintf("Geometry %T \narea= %f\nperimeter= %f\n\n", reflect.TypeOf(shape), area, perimeter))
	}
}
