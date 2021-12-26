package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここをカスタム
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	// i := []int{5, 1, 3, 3, 4, 5, 4, 7, 10}
	// s := []string{"z", "e", "a"}

	// fmt.Println(i, s)
	// sort.Ints(i)
	// sort.Strings(s)

	// fmt.Println(i, s)

	// el := []Entry{
	// 	{"A", 20},
	// 	{"F", 30},
	// 	{"Q", 90},
	// 	{"K", 50},
	// 	{"E", 10},
	// 	{"A", 30},
	// 	{"a", 80},
	// 	{"e", 30},
	// }

	// fmt.Println(el)
	// 第一引数にスライスを渡して、第二引数に条件式を渡す
	// sort.Slice(el, func(i, j int) bool {
	// 	// fmt.Println(i, j)
	// 	return el[i].Name < el[j].Name
	// })
	// fmt.Println("---------")
	// fmt.Println(el)
	// fmt.Println("---------")

	// sort.SliceStable(el, func(i, j int) bool {
	// 	return el[i].Name < el[j].Name
	// })
	// sort.SliceStable(el, func(i, j int) bool {
	// 	return el[i].Value < el[j].Value
	// })
	// fmt.Println("---------")
	// fmt.Println(el)
	// fmt.Println("---------")

	// カスタムソート
	m := map[string]int{"S": 1, "J": 2, "A": 3, "N": 3}
	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	fmt.Println(lt)

	sort.Sort(lt)

	fmt.Println(lt)
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}
