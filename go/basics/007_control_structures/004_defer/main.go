package main

import (
	"fmt"
	"os"
)

// defer
// 関数の終了時に実行される処理を登録できる

func TestDefer() {
	defer fmt.Println("End")
	fmt.Println("Start")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()
	// 出力:
	// Start
	// End

	// deferは関数の終了時に実行される
	// 関数の中でdeferを使うと、関数が終了する前に登録された処理が実行される
	// 例えば、ファイルのクローズ処理や、リソースの解放処理などに使われる

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
		// 1,2,3
	}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // ファイルを閉じる処理を登録

	file.Write([]byte("Hello, World!")) // ファイルに書き込む
}
