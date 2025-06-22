package main

import (
	"fmt"
	"strconv"
)

// if
// 条件分岐
// エラーハンドリング

func main() {
	a := 0
	if a == 2 {
		fmt.Println("2")
	} else if a == 1 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

	// 簡易文付きif文
	if b := 100; b == 100 {
		fmt.Println("b is 100")
	} else {
		fmt.Println("b is not 100")
	}

	// エラーハンドリング
	var s string = "AIUEO"
	i, err := strconv.Atoi(s) // 文字列を整数に変換
	if err != nil {
		fmt.Println("Error:", err) // Error: strconv.Atoi: parsing "AIUEO": invalid syntax
	}
	fmt.Println(i) // エラーがなければ変換結果を表示
}
