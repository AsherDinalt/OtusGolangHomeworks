package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	cr := Circle{10}
	rc := Rectangle{10, 5}
	tr := Triangle{8, 6}
	ar, ok := calculateArea(cr)
	if ok == nil {
		fmt.Println(ar)
	}
	ar, ok = calculateArea(rc)
	if ok == nil {
		fmt.Println(ar)
	}

	ar, ok = calculateArea(tr)
	if ok == nil {
		fmt.Println(ar)
	}
}

func calculateArea(s any) (float64, error) {
	sh, ok := s.(Shape)
	if !ok {
		return 0.0, errors.New("not Shape interface")
	}
	return sh.Area(), nil
}
