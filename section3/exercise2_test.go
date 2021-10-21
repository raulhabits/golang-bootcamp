package section3

import (
	"fmt"
	"testing"
)

func TestIntMinMultipleAsserts(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First value less than second, expected first", args{1, 2}, 1},
		{"Second value less than first, expected second", args{2, 1}, 1},
		{"First value equals than second, expected any", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}