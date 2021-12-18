package main

import "fmt"

// struct
// オブジェクト指向で言うクラスのようなもの

type User struct {
	Name string
	Age  int
	// まとめて定義も可能
	// X, Y int
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 19
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 100}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// Nameだけ値が入る、Ageは初期値
	user5 := User{Name: "user5"}
	fmt.Println(user5)

	// newをするとポインタ型として結果に&がつく
	user6 := new(User)
	fmt.Println(user6)

	// こっちのほうが多い
	user7 := &User{}
	fmt.Println(user7)

	UpdateUser(user1)
	fmt.Println(user1)

	UpdateUser2(user7)
	fmt.Println(user7)
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}
