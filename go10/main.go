package main

import (
	"fmt"
	"time"
)

// func main() {

// 	defer func() {
// 		if x := recover(); x != nil {
// 			fmt.Println(x)
// 		}
// 	}()

// 	panic("runtime error")
// 	fmt.Println("Start")
// }

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep((200 * time.Millisecond))
	}
}
