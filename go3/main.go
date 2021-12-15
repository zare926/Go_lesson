package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int型
	// var i int64 = 100

	// fmt.Printf("%T\n", i)
	// // キャスト
	// fmt.Printf("%T\n", int32(i))

	// float型　浮動小数点型
	var fl64 float64 = 2.4

	fmt.Println(fl64)

	// fl := 3.2 暗黙型の場合自動でfloat64になる

	var t, f bool = true, false
	fmt.Println(t, f)

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	// var arr1 [3]int

	// [...]で要素数を勝手に判定
	// 要素数はlengthじゃなくlen()

	// interface型　あらゆる型に対応する
	// 計算ができない
	var x interface{}
	fmt.Println(x)

	// fl := 10.5
	// i := int(fl)
	// fmt.Println(i)

	i, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

}
