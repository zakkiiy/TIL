package main

// 変数
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now())

	// 明示的な変数定義
	var i int = 100
	fmt.Println(i)

	var s string = "Hello, Go!"
	fmt.Println(s)

	// 複数の変数にまとめて定義可能（同じ型）
	var t, f bool = true, false
	fmt.Println(t, f)

	// 異なる型の変数をまとめて定義
	var (
		i2 int    = 1
		s2 string = "Hello, Go!"
	)
	fmt.Println(i2, s2)

	// 型のみ定義
	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 初期値は0と空文字列

	i3 = 10
	s3 = "Hello, World!"
	fmt.Println(i3, s3)

	// 暗黙的な型定義
	// 関数内でのみ定義できる
	i4 := 100
	fmt.Println(i4)

	// 再代入
	i4 = 200
	fmt.Println(i4)

	// 再定義はできない
	// i4 := 300
}
