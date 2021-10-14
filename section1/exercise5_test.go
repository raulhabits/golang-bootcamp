package section1

import (
	"fmt"
	"testing"
)

func Test_passedParametersShowcase(t *testing.T) {
	number := 10

	fmt.Println("Test_passedParametersShowcase Original number", number)

	parameterPassedPerValue(number)

	fmt.Println("Test_passedParametersShowcase after being finished the function passed Per Value", number)

	parameterPassedPerReference(&number)

	fmt.Println("Test_passedParametersShowcase  being finished the function passed Per Reference", number)

}
