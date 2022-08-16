package main

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() (perim float64) {
	perim = 2 * (r.Height + r.Width)
	return
}

func (r Rectangle) Area() (area float64) {
	area = (r.Height * r.Width)
	return
}

func (c Circle) Area() (area float64) {
	area = 3.14 * c.Radius * c.Radius
	return
}

func (c Circle) Perimeter() (perim float64) {
	perim = 2 * c.Radius * 3.14
	return
}
func (t Triangle) Area() (area float64) {
	area = t.Base * t.Height / 2
	return
}
func (t Triangle) Perimeter() (perim float64) {
	perim = 3 * t.Base
	return
	//Only works for equilateral triangles
}
