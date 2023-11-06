package main

import (
    "fmt"
)

// Shape interface
// Add your solution here!

func main() {
    circle := Circle{Radius: 5.0}
    rectangle := Rectangle{Width: 3.0, Height: 4.0}

    shapes := []Shape{circle, rectangle}

    for _, shape := range shapes {
        fmt.Println(shape.Name(), shape.Area())
    }
}
