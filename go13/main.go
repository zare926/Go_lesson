package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 590)
	fmt.Println(sl)

	sl = append(sl, 590, 500, 400, 300)
	fmt.Println(sl)

	sl2 := make([]int, 6)
	fmt.Println(sl2)
	fmt.Println(len(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	// COPY
	slcopy := []int{1, 2, 3, 4, 5}
	slcopy2 := make([]int, 5, 10)
	n := copy(slcopy2, slcopy) //nはコピーに成功した数を返す
	fmt.Println(n, slcopy2)
}
