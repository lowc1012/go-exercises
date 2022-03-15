// The Recruiter (James) told me that I could answer these questions 
// in Go or Python, so I deleled all the Java code and chose Go.
package main

import "fmt"

type Shape interface {
    getArea()      float64
    getPerimeter() float64
    calArea()      float64
    calPerimeter() float64
    toString()     string
}

type Rectangle struct {
    length float64
    width  float64
}

func NewRectangle(newLength, newWidth float64) *Rectangle {
    r := &Rectangle{
        length: newLength,
        width: newWidth,
    }
    return r
}

func (r Rectangle) getArea() float64 {
    s := "Finding area of rectangle with length = %.1f and width = %.1f\n"
    if r.length == r.width {
        s = "Finding area of Square with side = %.1f\n"
        fmt.Printf(s, r.length)
    } else {
        fmt.Printf(s, r.length, r.width)
    }
    return r.calArea()
}

func (r Rectangle) calArea() float64 {
    return r.length * r.width
}

func (r Rectangle) getPerimeter() float64 {
    
    s := "Finding perimeter of rectangle with length = %.1f and width = %.1f\n"
    if r.length == r.width {
        s = "Finding perimeter of Square with side = %.1f\n"
        fmt.Printf(s, r.length)
    } else {
        fmt.Printf(s, r.length, r.width)
    }

    return r.calPerimeter()
}

func (r *Rectangle) calPerimeter() float64 {
    if r.length == r.width {
        return 4 * r.length
    }
    return (r.length + r.width) * 2
}

func (r Rectangle) toString() string {
    s := "Rectangle = [length: %.1f, width: %.1f, area: %.1f, perimeter: %.1f]"
    if r.length == r.width {
        s = "Square = [side: %.1f, area: %.1f, perimeter: %.1f]"
        return fmt.Sprintf(
            s,
            r.length,
            r.calArea(),
            r.calPerimeter(),
        )
    }
    return fmt.Sprintf(
        s,
        r.length,
        r.width,
        r.calArea(),
        r.calPerimeter(),
    )
}

type Circle struct {
    radius float64
}

func NewCircle(newRadius float64) *Circle {
    c := &Circle{
        radius: newRadius,
    }
    return c
}

func (c Circle) getArea() float64 {
    fmt.Printf("Finding area of Circle with radius = %.1f\n", c.radius)
    return c.calArea()
}

func (c Circle) calArea() float64 {
    return 3.14 * c.radius * c.radius
}

func (c Circle) getPerimeter() float64 {
    fmt.Printf("Finding perimeter of Circle with radius = %.1f\n", c.radius)
    return c.calPerimeter()
}

func (c Circle) calPerimeter() float64 {
    return 6.28 * c.radius
}

func (c Circle) toString() string {
    return fmt.Sprintf(
        "Circle = [radius: %.1f, area: %.2f, perimeter: %.2f]",
        c.radius,
        c.calArea(),
        c.calPerimeter(),
    )
}

type Square struct {
    Rectangle
}

func NewSquare(side float64) *Square {
    return &Square{
        Rectangle{
            width: side,
            length: side,
        },
    }
}

func main() {
    
    var length float64
    var width float64
    var radius float64
    var side float64
    
    fmt.Println("New rectangle. Please input the length and width")
    fmt.Scanln(&length, &width)
    
    fmt.Println("New circle. Please input the radius")
    fmt.Scanln(&radius)

    fmt.Println("New square. Please input the side")
    fmt.Scanln(&side)

    rec := NewRectangle(length, width)
    fmt.Printf("Area = %.1f and Perimeter = %.1f\n", rec.getArea(), rec.getPerimeter())

    squ := NewSquare(side)
    fmt.Printf("Area = %.1f and Perimeter = %.1f\n", squ.getArea(), squ.getPerimeter())

    cir := NewCircle(radius)
    fmt.Printf("Area = %.2f and Perimeter = %.2f\n", cir.getArea(), cir.getPerimeter())
    
    fmt.Println(rec.toString())
    fmt.Println(squ.toString())
    fmt.Println(cir.toString())
}
