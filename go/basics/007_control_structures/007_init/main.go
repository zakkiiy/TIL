package main

import "fmt"

func init() {
	// init 関数はパッケージが初期化されるときに自動的に呼び出される
	// パッケージ内で一度だけ実行される
	// main パッケージでは、main 関数の前に実行される
	// 複数の init 関数がある場合、定義された順に実行される
	// 変数の初期化や設定などに使用される

	fmt.Println("init function called")
}

func main() {
	fmt.Println("main function called")
	// init 関数は main 関数の前に実行される
	// ここでは、init 関数の効果を確認するためにメッセージを表示
	// init 関数は自動的に呼び出されるため、明示的に呼び出す必要はない
}
