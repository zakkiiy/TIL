package main

import (
	"fmt"
	"strconv"
)

// 型

func main() {
	// int
	var i int = 100
	fmt.Println("i:", i)
	i = 200 // 再代入
	fmt.Println("i:", i)

	// %Tは型を表示するフォーマット指定子
	fmt.Printf("%T\n", i) // int

	// float
	var f64 float64 = 3.14
	fmt.Println("f:", f64)

	fl := 2.718 // float64になる 暗黙
	fmt.Printf("%T\n", fl) // float64

	// unit型, complex型
	// 使用頻度少ない

	// bool 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	// string 文字列型
	// byte配列の集まり
	var s string = "Hello, Go!"
	fmt.Println(s)

	fmt.Println(`test
	test
							test`)

	fmt.Println(s[0]) // byte型で表示される 72
	fmt.Println(string(s[0])) // string型に変換して表示される H

	// byte(unit8)型
	byteA := []byte{72,73}
	fmt.Println(byteA) // [72 73]

	fmt.Println(string(byteA)) // HI

	c := []byte("HI")
	fmt.Println(c) // [72 73]
	fmt.Println(string(c)) // HI

	// 配列
	// 後から要素数を増やせない
	var arr1 [3]int
	fmt.Println(arr1) // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 = [3]string{"A", "B", "C"}
	fmt.Println(arr2) // [A B C]
	fmt.Printf("%T\n", arr2) // [3]string}

	// 暗黙的
	arr3 := [3]string{"go", "ruby"}
	fmt.Println(arr3) // [go ruby ""]
	
	arr4 := [...]string{"D", "E"} // 要素数を自動でカウントする
	fmt.Println(arr4) // [D E]
	fmt.Printf("%T\n", arr4) // [2]string

	fmt.Println(arr2[0]) // A
	fmt.Println(arr2[2-1]) // B
	arr2[2] = "D"
	fmt.Println(arr2) // [A B D]

	// 要素の中身を変更することは可能だか、要素数以上の要素を追加することはできない 
	// arr2[3] = "E" // エラー: index out of range [3] with length 3

	// len 関数で要素数を取得できる
	fmt.Println(len(arr2)) // 3


	// interface型 & nil
	// あらゆる型を受け入れることができる
	var x interface{} // 空のinterface型
	fmt.Println(x) // <nil>
	x = 100 // int型を代入
	fmt.Println(x) // 100
	fmt.Printf("%T\n", x) // int
	x = "Hello, Go!" // string型を代入
	fmt.Println(x) // Hello, Go!
	x = []int{1, 2, 3} // スライス型を代入
	fmt.Println(x) // [1 2 3]
	fmt.Printf("%T\n", x)

	// x = 2
	// fmt.Println(x + 3) // エラー interface型とint型は直接演算できない
	
	// 型変換
	var i2 int = 1
	fl64 := float64(i2) // intからfloat64へ変換
	fmt.Printf("i2 = %T\n", i2) // int
	fmt.Printf("fl64 = %T\n", fl64) // float64
	fmt.Println(fl64) // 1

	// 文字列から数値への変換
	var s2 string = "123"
	// _ で戻り値を破棄できる
	i3, _ := strconv.Atoi(s2)
	fmt.Println(i3) // 123

	// 数値から文字列への変換
	var s3 int = 200
	i4 := strconv.Itoa(s3)
	fmt.Println(i4) // 200
	fmt.Printf("%T\n", i4) // string

	// 文字列をバイト配列に変換
	var h string = "Hello, Go!"
	b := []byte(h)
	fmt.Println(b) // [72 101 108 108 111 44 32 71 111,33]

	// バイトを文字列に変換
	h2 := string(b)
	fmt.Println(h2) // Hello, Go!
}


// 型を指定することで、誤った値の代入や処理を防げる
// コンパイル時にエラーを検出でき、開発中の安全性が高まる
// コード補完やリファクタリングもしやすくなる
