package main

import "fmt"

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	p2 := &Point{200, "ABCS"}
	fmt.Println(p)
	fmt.Println(p2)
}
