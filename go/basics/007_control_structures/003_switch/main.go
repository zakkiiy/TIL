package main

import "fmt"

// switch

// 型スイッチ
func anything(v interface{}) {
	fmt.Println(v)
}

func main() {
	n := 1
	switch n {
	case 1, 2:
		println("1 or 2")
	case 3:
		println("3")
	default:
		println("other")
	}

	switch n := 10; n {
	case 1, 2:
		fmt.Println("n is 1 or 2")
	case 3:
		fmt.Println("n is 3")
	default:
		fmt.Println("n is other")
	}

	// 型スイッチ
	var x interface{} = 3
	i, isInt := x.(int)     // 型アサーション
	fmt.Println(i+4, isInt) // 7

	// 様々な型に対応したinterface{}
	// interface{} 型に入っている値の型を判定して分岐する例
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Println("x is int:", i)
	} else if s, isString := x.(string); isString {
		fmt.Println("x is string:", s)
	} else {
		fmt.Println("x is of a different type")
	}

	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case bool:
		fmt.Println("x is bool")
	default:
		fmt.Println("x is of a different type")
	}
}
