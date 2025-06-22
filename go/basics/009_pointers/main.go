package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s ...int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)  //100
	fmt.Println(&n) // 0x14000104020

	Double(n)
	fmt.Println(n) // 100

	// ポインタ型で宣言
	// アドレスを値渡し
	// int型のポインタ
	// &n は：変数 n の アドレスを取得
	var p *int = &n
	// p は「int 型の値が格納されたアドレスを保持する変数」です
	// p 自体は、ただの アドレス（ポインタ） です
	fmt.Println(p)  // 0x1400000e0d0
	fmt.Println(*p) // 100

	// *p は「ポインタ p が指しているアドレスの中身（=n）」
	// だから *p = 300 は「n を 300 に書き換える」のと同じ
	*p = 300
	fmt.Println(n) // 300
	n = 500
	fmt.Println(*p) // 500
	// p はただの アドレス です（つまり「n が格納されている場所」）
	// なので fmt.Println(p) とすると、0x1400... のようなアドレス文字列が出力される。
	// *p はポインタ p が指している先（アドレスの中身）を参照する

	// nはこの時点で500
	// &n アドレスを渡している
	Double2(&n)
	fmt.Println(n) // 1000

	/*
		i *int：これは int型のポインタ を引数に取る、という意味
		*i = *i * 2：
		*i → アドレス先の中身（つまり、呼び出し元の変数の値）
		*i * 2 → その値を2倍
		*i = ... → その結果を元の場所（n）に書き戻す

		&n：nのアドレスを渡している（アドレスを値渡ししている）
		関数側では *i により、nの中身そのものを操作している
	*/

	// スライス
	// 参照型はもともとその機能を持っているためポインタ不要
	sl := []int{1, 2, 3, 4, 10}
	Double3(sl...)
	fmt.Println(sl) // [2 4 6]

}

// コピーしたnを関数の引数に渡している。
// 元のデータは変更されない
// ポインタを使用することで、解消可能

/*
func Double(i int) int {
	i = i * 2
	return i
}

func main() {
	var n int = 100
	fmt.Println(n)  //100
	fmt.Println(&n) // 0x14000104020

	n := Double(n)
	fmt.Println(n) // 100

}
*/
