package main

import "fmt"

// initのほうがmainより先に読まれる
// 初期化処理をしたい時に定義するケースが多い
// 2個書いた場合は上から順番に読まれるが、各ケースは殆どない。

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
