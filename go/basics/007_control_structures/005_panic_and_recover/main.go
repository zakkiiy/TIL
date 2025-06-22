package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // パニックから回復: runtime error
		}
	}()
	panic("runtime error")                        // パニックを発生させる panic: runtime error
	fmt.Println("This line will not be executed") // この行は実行されない
}

// defer は関数終了時に必ず実行される（panicがあっても）
// panic 時、defer 内で recover() を使うことで回復可能
// recoverはdefer 内でしか使用できない
// recover が nil 以外 → panic が発生していたことを示す

/*
panic() が発生すると：
	その時点で関数の処理は中断されます
	すぐに スタックを巻き戻しながら defer を順に実行
	recover() がなければ、プログラムはクラッシュして終了
	recover() があれば、panic は回収されて、処理は継続可能（ただし panic 後の処理には戻らない）
*/
