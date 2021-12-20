package main

import (
	"Go_lesson/go29/foo"
	// "fmt"
	. "fmt"
	f "fmt"
)

func Do(s string) (b string) {
	Println(b, "1")
	Println(s, "2")
	b = s
	{
		// b := "BBBB"
		// Println(b)
	}
	// Println(b)
	return b
}

func main() {
	f.Println(foo.Max)

	Println(foo.RetrunMin())

	Do("AAAA")
}
