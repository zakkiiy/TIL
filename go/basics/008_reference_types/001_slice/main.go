package main

import "fmt"

// スライス　可変長引数
func Sum(s ...int) int {
	fmt.Println(s) // [1 2 3 4 10]
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// スライス　可変長引数
	k := Sum(1, 2, 3, 4, 10)
	fmt.Println(k) // 20

	// スライスを渡すことも可能
	k2 := []int{1, 2, 3}
	fmt.Println(Sum(k2...)) // [1 2 3]

	var sl = []int{1, 2, 3, 4, 5}
	fmt.Println(sl) // [1 2 3 4 5]

	// 暗黙的な宣言
	sl2 := []string{"a", "b", "c"}
	fmt.Println(sl2) // [a b c]

	// make関数でスライスを作成
	sl3 := make([]int, 5)
	fmt.Println(sl3) // [0 0 0 0 0]

	// スライスの要素変更
	sl3[0] = 1000
	fmt.Println(sl3) // [1000 0 0 0 0]

	// 値の取り出し
	fmt.Println(sl3[0]) // 1000

	// スライス
	// append make len cap
	sl4 := []int{3, 4, 5}
	fmt.Println(sl4)

	// append
	sl4 = append(sl4, 6, 7, 8)
	fmt.Println(sl4) // [3 4 5 6 7 8]

	// make スライスを生成する関数
	sl5 := make([]int, 5)
	fmt.Println(sl5) // [0 0 0 0 0]

	// len
	fmt.Println(len(sl5)) // 5

	sl6 := make([]int, 5, 10)
	fmt.Println(len(sl6)) // 5
	// cap 容量
	fmt.Println(cap(sl6)) // 10

	// スライスfor
	sl7 := []string{"a", "b", "c"}
	for i := range sl7 {
		fmt.Println(i, sl[i])
	}
}

// append は新しいスライスを返す関数です（元のスライスを変更しない）。
// そのため、返り値を同じ変数 sl4 に代入するのは正しい使い方です。
