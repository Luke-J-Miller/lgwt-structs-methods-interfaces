package main

import "testing"

func TestPerimeter(t *testing.T) {

	perimTests := []struct {
		name              string
		shape             Shape
		expectedPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedPerimeter: 62.8},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedPerimeter: 36.0},
	}
	for _, tt := range perimTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got-tt.expectedPerimeter > .001 || got-tt.expectedPerimeter < -.001 {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.expectedPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedArea: 314.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got-tt.expectedArea > .001 || got-tt.expectedArea < -.001 {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.expectedArea)
			}
		})
	}
}
