package section1

import (
	"fmt"
	"testing"
)

func Test_passedParametersShowcase(t *testing.T) {
	number := 10

	fmt.Println("Test_passedParametersShowcase Original number", number)

	parameterPassedPerValue(number)

	fmt.Println("Test_passedParametersShowcase Passed Per Value", number)

	parameterPassedPerReference(&number)

	fmt.Println("Test_passedParametersShowcase Passed Per Reference", number)

}
