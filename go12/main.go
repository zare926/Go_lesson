package main

import "fmt"

func main() {
	var sl []int
	fmt.Println(sl)
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)
	var sl3 []string = []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 3)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[2])

	// 取り出し方2番から4番まで
	fmt.Println(sl5[2:4])
	// 2番まで
	fmt.Println(sl5[:2])
	// 2番から最後まで
	fmt.Println(sl5[2:])
	// 全部
	fmt.Println(sl5[:])
	// 最後
	fmt.Println(sl5[len(sl5)-1])
	//  最初と最後を抜く
	fmt.Println(sl5[1 : len(sl5)-1])

}
