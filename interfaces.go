//
// gobyexample.com
// interfaces.gobyexample
//
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	permi() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

funct (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permi())
}

func main() {
	r := rect{widt}
	c := circle(radius: 5}
		

	
}