package section1

import "fmt"

func parameterPassedPerValue(number int) {
	number++
	fmt.Println("Value on the parameterPassedPerValue function after increment ", number)
}

func parameterPassedPerReference(number *int) {
	*number++
	fmt.Println("Value on the parameterPassedPerReference function after increment ", number)
}