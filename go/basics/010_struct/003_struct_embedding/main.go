package main

import "fmt"

// 構造体の埋め込み
// フィールドに構造体を埋め込むことが可能

// T 経由で User のフィールドやメソッドにアクセスできる
// でも、User 型自体も変わらずそのまま単独で使える
type T struct {
	// User User 省略可能
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t) // {{user1 10}}

	// アクセス
	fmt.Println(t.User)      // {user1 10}
	fmt.Println(t.User.Name) // user1
	fmt.Println(t.Name)      // user1

	t.User.SetName()
	fmt.Println(t)
}

/*
	フィールド名	意味
	大文字始まり	外部パッケージからアクセス可能（エクスポートされる）
	小文字始まり	外部からは見えない（パッケージ内限定）
*/
