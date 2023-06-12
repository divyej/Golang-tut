package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    rectangle := Rectangle{
        Width:  5,
        Height: 3,
    }
    PrintArea(rectangle)
}
