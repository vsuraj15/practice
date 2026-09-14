package main

import "fmt"

type Measurement interface {
	Area()
	Perimeter()
}

type Square struct {
	a int
	b int
}

type Rectangle struct {
	x int
	y int
}

func (s Square) Area() {
	fmt.Printf("Area calulated: %+v\n", s.a*s.b)
}
func (s Square) Perimeter() {
	fmt.Printf("Perimeter calulated: %+v\n", 2*(s.a+s.b))
}

func (s Rectangle) Area() {
	fmt.Printf("Area calulated: %+v\n", s.x*s.y)
}
func (s Rectangle) Perimeter() {
	fmt.Printf("Perimeter calulated: %+v\n", 2*(s.x+s.y))
}

func main() {
	s := Square{a: 2, b: 2}
	r := Rectangle{x: 2, y: 5}
	calculateMeasurement(s)
	calculateMeasurement(r)
}

func calculateMeasurement(val interface{}) {
	switch v := val.(type) {
	case Square:
		fmt.Println("Received Square")
		v.Area()
		v.Perimeter()
	case Rectangle:
		fmt.Println("Received Rectangle")
		v.Area()
		v.Perimeter()
	default:
		fmt.Println("Unexpected type received")
	}
}
