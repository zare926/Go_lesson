package main

import "fmt"

// 頭文字を大文字にすると他のパッケージからも呼び出せる。
const Pi = 3.14
const (
	URL      = "http"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c1 = iota
	c2
	c3
)

func main() {
	fmt.Println(Pi)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(c1, c2, c3)
}
