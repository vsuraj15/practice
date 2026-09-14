package main

import "fmt"

type Measurement interface {
	Area()
	Perimeter()
}

type Square struct {
	a, b int
}

type Rectangle struct {
	a, b int
}

func (s Square) Area() {
	fmt.Printf("Area: %+v\n", s.a*s.b)
}

func (s Square) Perimeter() {
	fmt.Printf("Perimeter: %+v\n", 2*(s.a+s.b))
}

func (s Rectangle) Area() {
	fmt.Printf("Area: %+v\n", s.a*s.b)
}

func (s Rectangle) Perimeter() {
	fmt.Printf("Perimeter: %+v\n", 2*(s.a+s.b))
}

func main() {
	s := Square{a: 10, b: 20}
	r := Rectangle{a: 2, b: 3}
	calculateMeasurement(s)
	calculateMeasurement(r)
}

func calculateMeasurement(v interface{}) {
	switch a := v.(type) {
	case Square:
		fmt.Println("Measurement for square")
		a.Area()
		a.Perimeter()
	case Rectangle:
		fmt.Println("Measurement for rectangle")
		a.Area()
		a.Perimeter()
	default:
		fmt.Errorf("Unexpected type found")
	}
}
