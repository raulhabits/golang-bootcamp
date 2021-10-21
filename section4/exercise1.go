package section4

import (
	"fmt"
	"time"
)
func GoRoutineForExecutingAnonymousPrintFunction(parameters interface{})  {
	go func(params interface{}) {
		fmt.Println(params)
	}(parameters)
	time.Sleep(time.Second)
}