package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (float64, int) {
	var q float64 = float64(x) / float64(y)
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noretrun() {
	fmt.Println("NO Return")
	return
}

func main() {
	i := Plus(10, 58)
	fmt.Println(i)

	i2, i3 := Div(55, 100)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noretrun()

	// 無名関数
	f := func(x, y int) int {
		return x * y
	}

	i5 := f(20, 40)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x * y
	}(100, 289)

	fmt.Println(i6)
}
