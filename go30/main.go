package main

import (
	"Go_lesson/go30/alib"
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(0))
	fmt.Println(IsOne(1))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
