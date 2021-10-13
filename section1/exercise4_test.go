package section1

import (
	"reflect"
	"testing"
)


func simpleQueueOnChannelWithData(input int) SimpleQueue {
	queue := SimpleQueue{}
	queue.init()
	queue.data <- input
	return queue
}


func TestSimpleQueue_dequeue(t *testing.T) {
	tests := []struct {
		name   string
		fields SimpleQueue
		want   interface{}
	}{
		{"Dequeue with existing elements", simpleQueueOnChannelWithData(1), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.fields
			if got := receiver.dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleQueue_enqueue(t *testing.T) {
	tests := []struct {
		name   string
		enqueue int
		expected int
	}{
		{"Enqueue success", 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := SimpleQueue{}
			receiver.init()
			receiver.enqueue(tt.enqueue)
			got := receiver.dequeue()
			if got != tt.expected {
				t.Errorf("Enqueue() = %v, want %v", got, tt.expected)
			}
		})
	}
}
