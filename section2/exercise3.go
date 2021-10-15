package section2

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (receiver *Circle) Area() float64 {
	return math.Pi * math.Exp2(receiver.Radius)
}

func (receiver *Circle) Perimeter() float64 {
	return 2 * math.Pi * receiver.Radius
}

func (receiver *Rectangle) Area() float64 {
	return receiver.Width * receiver.Height
}

func (receiver *Rectangle) Perimeter() float64 {
	return 2 * receiver.Width * receiver.Height
}

func performAreaAndPerimeter(shape Shape) (float64, float64) {
	return shape.Area(), shape.Perimeter()
}
