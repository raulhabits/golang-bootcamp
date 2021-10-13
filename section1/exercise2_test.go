package section1

import (
	"testing"
)

func TestCalculator_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Add method test", args{1, 2}, 3},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator := &Calculator{}
			if got := calculator.add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		expectedPanic bool
		want int
	}{
		{ "Divide with valid parameters", args{1, 2}, false, 0},
		{ "Divide with non-valid parameters", args{1, 0}, true, 0},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tt.expectedPanic {
					t.Errorf("Expected panic!!!!!!")
				}

			}()
			calculator := &Calculator{}
			if got := calculator.divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_multiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Multiply method test", args{1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator := &Calculator{}
			if got := calculator.multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_subtract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Subtract method test", args{1, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator := &Calculator{}
			if got := calculator.subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
