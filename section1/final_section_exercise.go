package section1

type MyStack struct {
	data []interface{}
}

func (receiver *MyStack) push(element interface{})  {
	receiver.data = append(receiver.data, element)
}

func (receiver *MyStack) pop() interface{} {
	defer func() {

	}()
	length := len(receiver.data)
	element := receiver.data[length - 1]
	length--
	receiver.data = receiver.data[:length]
	return element
}
func (receiver *MyStack) peek() interface{} {
	defer func() {

	}()
	length := len(receiver.data)
	element := receiver.data[length - 1]
	return element
}
