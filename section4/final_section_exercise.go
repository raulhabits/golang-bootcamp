package section4

import (
	"fmt"
	"runtime"
	"strings"
)



func processElement(elementsChannel *chan string, element, search string)  {
	if strings.Contains(element, search) {
		*elementsChannel <- element
	}
}

func ArrayFilterBySubstring(elements []string, search string) {
	var numCPU = runtime.NumCPU()

	arraySize := len(elements)
	chunkSize := (arraySize + numCPU - 1) / numCPU

	var matchElements = make(chan string, 1)

	for i := 0; i < arraySize; i += chunkSize {
		end := i + chunkSize
		for _, elementOnArray := range elements[i:end] {
			go processElement(&matchElements, elementOnArray, search)
		}
	}

	for {
		select {
		case item := <- matchElements:
			fmt.Println(item)
		default:
			return
		}
	}
}