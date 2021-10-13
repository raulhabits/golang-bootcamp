package section1

import "errors"

var Products = map[string]*Product{}

type Product struct {
	ID string
	Name string
}

func Add(product *Product) error {
	if product.ID == "" {
		return errors.New("empty product id")
	}
	if Products[product.ID] != nil {
		return errors.New("product already exists")
	}
	Products[product.ID] = product
	return nil
}