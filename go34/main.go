package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"

}

func main() {
	ch := make(chan string)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("#######Error#######")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}
}
