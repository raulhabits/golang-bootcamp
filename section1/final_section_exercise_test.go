package section1

import (
	"reflect"
	"testing"
)

func prepareExistingData (values []int) []interface{}{
	var anything []interface{}
	for _, value := range values {
		anything = append(anything, value)
	}
	return anything
}

func TestMyStack_peek(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		errorExpected bool
		want   interface{}
	}{
		{"Empty stack", fields{}, true, 0},
		{"Empty stack", fields{prepareExistingData([]int {1, 2, 1000})}, true, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tt.errorExpected {
					t.Errorf("Expected panic!!!!!!")
				}

			}()
			receiver := &MyStack{
				data: tt.fields.data,
			}
			if got := receiver.peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peek() = %v, want %v", got, tt.want)
			}
			if got := receiver.peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peek() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestMyStack_pop(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		errorExpected bool
		want   []interface{}
	}{
		{"Empty stack", fields{}, true, prepareExistingData([]int{})},
		{"Pop value two times", fields{prepareExistingData([]int {1, 2, 1000})}, true, prepareExistingData([]int{1000, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tt.errorExpected {
					t.Errorf("Expected panic!!!!!!")
				}

			}()
			receiver := &MyStack{
				data: tt.fields.data,
			}
			for _, value := range  tt.want {
				if got := receiver.pop(); !reflect.DeepEqual(got, value) {
					t.Errorf("peek() = %v, want %v", got, value)
				}
			}
		})
	}
}


func TestMyStack_push(t *testing.T) {

	tests := []struct {
		name   string
		data []interface{}
		errorExpected bool
		want   []interface{}
	}{
		{"Multiple Push values", prepareExistingData([]int {1, 2, 1000}), true, prepareExistingData([]int{1000, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tt.errorExpected {
					t.Errorf("Expected panic!!!!!!")
				}

			}()
			receiver := &MyStack{}

			for _, value := range tt.data {
				receiver.push(value)
			}

			for _, value := range  tt.want {
				if got := receiver.pop(); !reflect.DeepEqual(got, value) {
					t.Errorf("peek() = %v, want %v", got, value)
				}
			}
		})
	}
}
