package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println(i) // 1 2
	}

	// 条件付き
	point := 0
	for point < 10 {
		fmt.Println(point) // 0 1 2 3 4 5 6 7 8 9
		point++
	}

	// 古典的
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // 3はスキップ
		}
		fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9
	}

	// 配列でfor
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i]) // 0 1 ... 4 5
	}

	// rangeを使ったfor
	arr2 := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr2 {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Value: %d\n", i, v) // Index: 0, Value: 1 ... Index: 4, Value: 5
	}

	// sliceを使ったfor
	// range sl でスライス sl を 1つずつ要素に分解してループ
	sl := []string{"apple", "banana", "cherry"}
	for i, v := range sl {
		fmt.Println("Index:", i, "Value:", v) // Index: 0, Value: apple ... Index: 2, Value: cherry
	}

	// mapを使ったfor
	// key, value形式　rubyのhash的な
	// keyが[string] valueがint
	m := map[string]int{"apple": 100, "banana": 200, "cherry": 300}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v) // Key: apple Value: 1 ... Key: cherry Value: 3
	}
}
