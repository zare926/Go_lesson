package main

import "fmt"

func main() {
	arr := [3]int{2, 4, 6}
	// iがindex番号になる
	// _にすれば要素だけ取れる
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
