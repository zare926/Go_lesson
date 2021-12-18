package main

import "fmt"

func main() {
	var n int = 100

	// ポイント型は&をつけると確認することができる
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	// 実態を表示するには*
	fmt.Println(*p)

	Double(&n)
	fmt.Println(n)
	Double(p)
	fmt.Println(n)
}

func Double(i *int) {
	*i = *i * 2
}
