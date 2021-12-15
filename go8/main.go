package main

import "fmt"

func anything(a interface{}) {
	// fmt.Println(a)

	switch v := a.(type) {
	case string:
		fmt.Println(v, "!?")
	case int:
		fmt.Println(v + 1000)
	}
}

func main() {

	anything("aiueo")
	anything(190)

	var x interface{} = 3
	f, isFloat64 := x.(float64)

	fmt.Println(f, isFloat64)

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
