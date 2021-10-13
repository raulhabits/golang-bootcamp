package section1

type SimpleQueue struct {
	data chan interface{}
}

func (receiver *SimpleQueue) init()  {
	receiver.data = make(chan interface{}, 1)
}

func (receiver *SimpleQueue) enqueue(element interface{})  {
	receiver.data <- element
}

func (receiver *SimpleQueue) dequeue() interface{} {
	return <- receiver.data
}

