package section3

import "testing"

func TestIntMin2(t *testing.T) {
	expected := 10
	min := 10
	max := 20
	result := IntMin(min, max)
	if result != expected {
		t.Errorf("IntMin() = %v, want %v", result, expected)
	}
}
