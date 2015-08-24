package main

import (
	"fmt"
	"math"
)

func main() {
	c := new(Circle)
	fmt.Println(c.x, c.y, c.r)
	c1 := Circle{x: 1, y: 1, r: 5}
	fmt.Println(c1.x, c1.y, c1.r)
	c2 := Circle{10, 10, 5} // if you know the order the fields are defined in
	fmt.Println(c2.x, c2.y, c2.r)

	fmt.Println("\nCircle area:", circleArea(&c1))
	fmt.Println("\nCircle area method:", c1.area())

	r := Rectangle{0, 10, 0, 10}
	fmt.Println("\nRectangle area method:", r.area(), "\n")

	p := Person{"Steve"}
	a := Android{p, "humanoid"}
	a.Person.Talk()
	a.Talk()

	fmt.Println("\nTotal area of circle and rectangle:", totalArea(&c1, &r))
	shapes := []Shape{c1, r}
	m := MultiShape{shapes}
	fmt.Println("\nMultiShape area:", m.area())
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// defaults for types:
//
// int : 0
// float : 0.0
// string : ""
// pointer : nil

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) area() float64 { //area() is the "receiver"
	// allows us to call the function with the . operator
	// known as a "method" now
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

// func (r *Rectangle) area() float64 {
// ^^ receives a pointer (*Rectangle)
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person        //embedded
	Model  string // not embedded
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
