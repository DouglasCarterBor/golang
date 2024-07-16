package forms

type Form interface {
	area() float64
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.1416 * c.radius * c.radius
}
