package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)

	logger := log.New(os.Stdout, "test", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")
}
