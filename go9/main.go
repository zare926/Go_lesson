package main

import "fmt"

func main() {

	// ラベル付きfor

	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理しない")
	// 		}
	// 		fmt.Println("処理しない")
	// 	}
	// 	fmt.Println("END")

	// Loop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 3; j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i, j, i*j)
	// 		}
	// 		fmt.Println("処理しない")
	// 	}

	// defer

	testDefer()
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// 	fmt.Println("4")
	// }()

	RunDefer()
}

func testDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}
