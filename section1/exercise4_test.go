package section1

import (
	"testing"
)

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
			receiver := &SimpleQueue{}
			receiver.init()
			receiver.enqueue(tt.enqueue)
			got := receiver.dequeue()
			if got != tt.expected {
				t.Errorf("Enqueue() = %v, want %v", got, tt.expected)
			}
		})
	}
}
