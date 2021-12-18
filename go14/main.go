package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	fmt.Println(Sum(1, 2, 3))
	sl2 := []int{1, 2, 3}
	fmt.Println(Sum(sl2...))
}

func Sum(s ...int) int {
	n := 1
	for _, v := range s {
		n += v
	}
	return n
}
