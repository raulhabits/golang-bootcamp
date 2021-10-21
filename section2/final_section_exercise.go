package section2

import "fmt"

type Product struct {
	Name string
	Price float64
}

type Book struct {
	Product
}

type Game struct {
	Product
}

type ProductDiscount interface {
	applyDiscount()
}

func (receiver Product) printProductData()  {
	fmt.Println(receiver.Name, receiver.Price)
}

func (receiver Book) applyDiscount() float64 {
	discountRate := 0.1
	return receiver.Price*(1 - discountRate)
}

func (receiver Game) applyDiscount() float64 {
	discountRate := 0.2
	return receiver.Price*(1 - discountRate)
}