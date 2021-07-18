package utils

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

//Perimeter will calculate perimeter of rectangle
func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.width + rectangle.height)
}

//Area will calculate area of rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}
