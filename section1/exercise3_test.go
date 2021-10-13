package section1

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		product *Product
	}
	tests := []struct {
		name    string
		args    args
		products map[string]*Product
		wantErr bool
	}{
		{ "Expected error on empty ID", args{&Product{Name: "IPHONE"}}, map[string]*Product{"12345": &Product{}},true},
		{ "Expected error on duplicated product ID", args{&Product{ID: "12345", Name: "IPHONE"}}, map[string]*Product{"12345": &Product{}},true},
		{ "Add success on non-empty list", args{&Product{ID: "123456", Name: "IPHONE"}}, map[string]*Product{"12345": &Product{}},false},
		{ "Add success on empty list", args{&Product{ID: "12345", Name: "IPHONE"}}, map[string]*Product{},false},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Products = tt.products
			if err := Add(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
