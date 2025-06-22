package main

import "fmt"

// map
// key, value形式

func main() {
	// 明示的
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) // map[A:100 B:200]

	// 暗黙的
	m2 := map[string]string{"A": "apple", "B": "banana"}
	fmt.Println(m2) // map[A:apple B:banana]

	// キーをint
	m3 := map[int]string{1: "python", 2: "go", 3: "ruby"}
	fmt.Println(m3) // map[1:python 2:go 3:ruby]

	// makeでmap生成
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	// 追加
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4) // map[1:JAPAN 2:USA]

	fmt.Println(m4[1]) // JAPAN

	// mapのエラーハンドリング
	// Go の map は、**存在しないキーにアクセスしても「エラーにならず、ゼロ値（初期値）を返す」**という仕様です。
	s, ok := m4[3]
	if !ok {
		fmt.Println("値なし")
	}
	fmt.Println(s, ok)

	// delete関数で削除
	delete(m4, 2)
	fmt.Println(m4) // map[1:JAPAN]

	// len関数
	fmt.Println(len(m4)) // 1
}

// append はスライス専用の操作で、要素の「順序」がある構造に対して要素を末尾に追加するものです。
// 要素の追加・更新は m[key] = value の形で行う
// make(map[KeyType]ValueType) で空の map を作成