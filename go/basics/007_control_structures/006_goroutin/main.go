package main

import (
	"fmt"
	"time"
)

// goroutine
// 簡単に並行処理が実装可能

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond) // 100ミリ秒待機
	}
}

func main() {
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond) // 200ミリ秒待機
	}
}
