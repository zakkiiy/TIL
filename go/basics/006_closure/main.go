package main

import "fmt"

// クロージャー

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーター
// ジェネレーターは、値を生成する関数で、状態を保持しながら次の値を生成する
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))         // ""
	fmt.Println(f("World"))         // "Hello"
	fmt.Println(f("Go"))            // "World"
	fmt.Println(f("!"))             // "Go"
	fmt.Println(f(""))              // "!"
	fmt.Println(f("Hello, World!")) // ""
	fmt.Println(f("Hello, Go!"))    // "Hello, World!"

	// ジェネレータ
	f2 := integers()
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 2
	fmt.Println(f2()) // 3
	// integers は 高階関数（関数を返す関数）
	// 返す関数は クロージャ（外部の変数 i を覚えている）
	// 結果として、状態を持った関数オブジェクトを作り出している

	// リクエストごとのユニークID発行
	// テストデータの連番生成
	// 繰り返し処理における「状態付きの処理ステップ」

	// クロージャー間で変数の共有はしてない

}

// Later関数は、無名関数（クロージャ）を返す関数
// Later() を呼んだ時点では、store という変数が空で初期化されるだけで、関数本体は実行されない
// f("hello") で、Later()は実行されず、戻り値として指定された無名関数が実行される

/*
関数の中で関数（無名関数など）を定義し、
外側の関数で定義した変数に中の関数からアクセスして使うことで、
その変数は「記憶される＝保持される」ようになる。
*/
