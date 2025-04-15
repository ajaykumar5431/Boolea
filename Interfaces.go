package array

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Circle) Area(radius float64) float64 {
	Multi := r.radius * r.radius
	return Multi
}

func (r Rectangle) Ar() {
	rec := r.Width * r.Height

	fmt.Println(rec)
}
