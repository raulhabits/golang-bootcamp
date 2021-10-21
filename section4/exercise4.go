package section4

import (
	"fmt"
	"sync"
	"time"
)

var Counter = 0
var m sync.Mutex

func increase()  {
	m.Lock()
	Counter++
	m.Unlock()
}

func ExecuteIncreaseGoRoutine()  {

	for i := 0 ; i < 10000; i++ {
		go increase()
	}

	time.Sleep(time.Second)

	fmt.Println(Counter)
}
