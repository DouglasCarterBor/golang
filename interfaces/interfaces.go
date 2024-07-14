package main

import "fmt"

type form interface {
	area() float64
}

func writeArea(f form) {
	fmt.Printf("Area: %0.2f\n", f.area())
}

type rectangle struct {
	height float64
	width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.1416 * c.radius * c.radius
}

func main() {
	fmt.Println("Interfaces...")
	r := rectangle{height: 10, width: 5}
	writeArea(r)

	c := circle{radius: 5}
	writeArea(c)

}
