package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello1")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello(&wg)

	wg.Add(1)
	go func() {
		wg.Done()
		fmt.Println("Hello2")
	}()

	wg.Wait()
}
