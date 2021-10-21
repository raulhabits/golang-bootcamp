package section4

import "fmt"

var ObjectChannel = make(chan interface{})

func AddChannelDataUsingGoRoutine(data interface{}) {
	go func(dataOnGoRoutine interface{}) {
	ObjectChannel <- dataOnGoRoutine
	}(data)

	fmt.Println(<- ObjectChannel)
}