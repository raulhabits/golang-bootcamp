package section4

import (
	"fmt"
	"time"
)

func GoRoutineForExecutingPrintFunction(goRoutinePrintData interface{}, mainFunctionPrintData interface{}, sleepSeconds int64)  {
	go fmt.Println(goRoutinePrintData)

	fmt.Println(mainFunctionPrintData)

	time.Sleep(time.Second * time.Duration(sleepSeconds))
}