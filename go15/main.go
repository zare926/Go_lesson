package main

import "fmt"

func main() {
	var m = map[string]int{
		"A": 100,
		"B": 200,
	}
	fmt.Println(m)

	m2 := map[string]int{
		"A": 100,
		"B": 200,
	}
	fmt.Println(m2)

	m3 := make(map[int]int)
	fmt.Println(m3)

	m3[0] = 1
	m3[1] = 10
	fmt.Println(m3)
	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "A"
	fmt.Println(m4)

	s, ok := m4[1] // 第2引数は取り出しに成功したかどうかの判定を返す
	fmt.Println(s, ok)

	s2, ok2 := m4[2]
	fmt.Println(s2, ok2)
	if !ok2 {
		fmt.Println("error")
	}

	m4[2] = "B"
	m4[3] = "C"
	fmt.Println(m4)

	delete(m4, 2)
	fmt.Println(m4)

	// for
	ex := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for i, v := range ex {
		fmt.Println(i, v)
	}
}
