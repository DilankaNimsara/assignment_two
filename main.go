package main

import "fmt"

type shape interface {
	getAria() float64
}

type triangle struct {
	height float64
	base   float64
}
type squire struct {
	sideLength float64
}

func main() {
	t := triangle{base: 6, height: 4}
	s := squire{sideLength: 4}
	printAria(t)
	printAria(s)
}

func printAria(sp shape) {
	fmt.Println(sp.getAria())
}

func (t triangle) getAria() float64 {
	return 0.5 * t.base * t.height
}

func (s squire) getAria() float64 {
	return s.sideLength * s.sideLength
}
