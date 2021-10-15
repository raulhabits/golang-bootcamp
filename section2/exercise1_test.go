package section2

import "testing"

func Test_printSpecifiedInterfaceParameterWithCustomMessagePerType(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Input type integer", args{1}, "Input is Integer value and is 1"},
		{"Input type string", args{"hello world"}, "Input is String value and is hello world"},
		{"Input type boolean", args{true}, "Input is Boolean value and is true"},
		{"Input type args", args{args{}}, "Unhandled data type section2.args"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printSpecifiedInterfaceParameterWithCustomMessagePerType(tt.args.input); got != tt.want {
				t.Errorf("printSpecifiedInterfaceParameterWithCustomMessagePerType() = %v, want %v", got, tt.want)
			}
		})
	}
}
