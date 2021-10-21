package section3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_filterByAgeRange(t *testing.T) {
	type args struct {
		minAge int
		maxAge int
		ages []int
	}
	tests := []struct {
		name string
		args args
		expected []int
	}{
		{"Input ages is an empty list", args{0,0,[]int{}}, []int{}},

		{"Filter applied with values on the ages array range", args{10,20,[]int{1, 2, 10, 15, 20, 25}}, []int{10, 15, 20}},

		{"Filter out of the ages array range", args{30,40,[]int{1, 2, 10, 15, 20, 25}}, []int{}},

		{"Filter with impossible match", args{40,30,[]int{1, 2, 10, 15, 20, 25}}, []int{}},

	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := filterByAgeRange(testCase.args.minAge, testCase.args.maxAge, testCase.args.ages)
			assert.ObjectsAreEqualValues(got, testCase.expected)
		})
	}
}
