package section1

import "fmt"

type Calculator struct  {}

func (calculator *Calculator) add(a, b int) int {
	return a+b
}
func (calculator *Calculator) subtract(a, b int) int {
	return a-b
}
func (calculator *Calculator) multiply(a, b int) int {
	return a*b
}
func (calculator *Calculator) divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			errorMessage := "Error trying to divide by zero value"
			fmt.Println(errorMessage)
			panic(errorMessage)
		}

	}()
	return a/b
}
