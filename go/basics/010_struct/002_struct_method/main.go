package main

import "fmt"

// 構造体のメソッド
// 任意の型に適したメソッド

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

// uには、User型のオブジェクト（user1）を受け取る
// uを使用することで、userオブジェクトのコピーを操作可能
// 今回は引数を使って、Nameフィールドを更新しようとしているが、uは値渡し（コピー）のためuser1本体には反映されない
// レシーバuはuser1のコピーなので、uの変更は元のuser1には反映されない
func (u User) SetName(name string) {
	u.Name = name
}

// ポインタ
// レシーバは原則ポインタ型にする
// u *User：これは「ポインタレシーバ」
func (u *User) setName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("user2")
	user1.SayName()

	// メソッド呼び出し時に &user1 になって、uにはポインタが渡ってくる
	// Goは、値型の変数からポインタレシーバのメソッドを呼ぶと、自動的にアドレス（&変数）を渡してくれます。
	// ポインタ → そのまま渡る
	// 値 → Goが &値 にしてくれる
	user1.setName2("user3")
	user1.SayName()

	// ポインタ自体を渡すことも可能
	user2 := &User{Name: "user4"}
	user2.setName2("B")
	user2.SayName()

}

// 関数：構造体とは独立して定義・呼び出す
// メソッド：構造体に紐づいた関数（レシーバを使って定義）
// user1.SayName() のようにオブジェクト指向的に呼び出せる
// メソッドにすると、その型に関係する処理をまとめられて可読性・再利用性が上がる

// 構造体からオブジェクトを生成して、そのオブジェクトの操作メインの関数を作るなら、メソッド化したほうがいい
// 構造体の状態（フィールド）を扱う操作はメソッド化すると意図が明確になりやすい
// メソッドにすれば user.DoSomething() のような自然な書き方ができ、可読性や保守性が上がる
// 単なるユーティリティ処理なら関数のままでもOK（構造体と直接関係ない場合）
