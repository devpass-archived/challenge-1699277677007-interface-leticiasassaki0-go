package main

import (
	"testing"
)

func TestShape(t *testing.T) {
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 3.0, Height: 4.0}

	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		name := shape.Name()
		if name != "Circle" && name != "Rectangle" {
			t.Errorf("Expected name to be 'Circle' or 'Rectangle', but got %s", name)
		}

		area := shape.Area()
		if area <= 0 {
			t.Errorf("Expected positive area, but got %f", area)
		}
	}
}
