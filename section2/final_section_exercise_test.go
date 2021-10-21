package section2

import (
	"fmt"
)

func ExampleBook_apply() {
	book := Book{
		Product{
			Name: "Risk",
			Price: 1000,
		},
	}
	book.printProductData()
	fmt.Println(book.applyDiscount())
	// Output:
	// Risk 1000
	// 900
}

func ExampleGame_applyDiscount() {
	game := Game{
		Product{
			Name: "Risk",
			Price: 1000,
		},
	}
	game.printProductData()
	fmt.Println(game.applyDiscount())
	// Output:
	// Risk 1000
	// 800
}

func ExampleProduct_printProductData() {
	product := Product{
		Name:  "IPHONE",
		Price: 10000,
	}
	product.printProductData()
	// Output:
	// IPHONE 10000
}
