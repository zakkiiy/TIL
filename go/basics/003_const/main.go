package main

import "fmt"

// 定数
// 定数は基本的に関数外（パッケージスコープ）に定義する
// 他の関数からも使いやすく、コードの意図が明確になる
// 関数内でしか使わないならローカルに定義してもOK

const Pi = 3.14
const (
	// 複数の定数をまとめて定義
	URL = "https://example.com"
	Port = 8080
)

const (
	c0 = iota // 0
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	// Pi = 3.14159 // エラー: 定数は再代入できない

	fmt.Println(URL, Port)

	fmt.Println(c0, c1, c2) // 0 1 2
}

// itoaは定数の自動インクリメントを行う
// 連番ジェネレータ
// Goに enum はないが iota で代用