package section2

import "fmt"

func printSpecifiedInterfaceParameterWithCustomMessagePerType(input interface{}) string{
	var message string
	switch input.(type) {
	case int:
		message = fmt.Sprintf("Input is Integer value and is %d",  input)
	case string:
		message = fmt.Sprintf("Input is String value and is %s",  input)
	case bool:
		message = fmt.Sprintf("Input is Boolean value and is %t",  input)
	default:
		message = fmt.Sprintf("Unhandled data type %T",  input)
	}

	fmt.Println(message)

	return message
	
}
