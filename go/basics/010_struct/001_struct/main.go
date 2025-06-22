package main

import "fmt"

// 構造体
// 複数の任意の型の値をまとめたもの

type User struct {
	// フィールド
	Name string
	Age  int
	// X, Y int 複数定義
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "B"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1) // { 0}

	// 各フィールドを更新
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	// 暗黙的な定義
	user2 := User{}
	fmt.Println(user2) // { 0}

	user3 := User{Name: "user3", Age: 20}
	fmt.Println(user3) // {user3 20}

	// 値だけ指定 順番決まっている
	user4 := User{"user4", 40}
	fmt.Println(user4) // {user4 40}

	user5 := User{Name: "user5"}
	fmt.Println(user5) // {user5 0}

	user7 := new(User)
	fmt.Println(user7)  // &{ 0}
	fmt.Println(&user7) // 0x1400010c038
	fmt.Println(*user7) // { 0}

	user8 := &User{}
	fmt.Println(user8) // &{ 0}

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

	// 後からポインタ化も可能
	// u := User{Name: "Taro", Age: 20}
	// p := &u
}
