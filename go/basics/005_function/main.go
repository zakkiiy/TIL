package main

import "fmt"

// 関数
// 引数と戻り値には必ず型が必要
func Plus(x int, y int) int {
	// (x, y int) 同じ型の場合、省略可能
	return x + y
}

// 複数の戻り値
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 戻り値の変数を指定することができる
// 名前つき戻り値　あまり使用されてないぽい？
func Double(price int) (result int) {
	result = price * 2
	return // 戻り値の変数名を指定しているので、returnだけでOK
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("Hello from returned function!")
	}
}

// 関数を引数に取る関数
func CallFUnc(f func()) {
	f() // 引数の関数を実行
}

// クロージャー

func main() {
	i := Plus(1, 2)
	println(i) // 3

	i2, i3 := Div(9, 3)
	// 戻り値は必ず使用する必要があるが、破棄もできる
	// i2, _ := Div(9, 3) // 破棄
	fmt.Println(i2, i3)

	i4 := Double(100)
	fmt.Println(i4) // 200

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}

	i5 := f(3, 4)
	fmt.Println(i5) // 7

	// 即時関数（IIFE: Immediately Invoked Function Expression）
	f2 := func(x, y int) int {
		return x * y
	}(5, 7)

	fmt.Println(f2) // 35

	// 関数を返す関数
	// f3に無名関数が代入される
	// f3は「関数」そのものになるので、あとで f3() として実行できる
	f3 := ReturnFunc()
	f3() // Hello from returned function!

	// 関数を引数に取る関数
	CallFUnc(func() {
		fmt.Println("I'm a function")
	})

}

// 4. 可変長引数もある！
// Yes！Goでも 可変長引数（variadic arguments） を使えます！

// go
// コピーする
// 編集する
// func Sum(nums ...int) int {
// 	sum := 0
// 	for _, n := range nums {
// 		sum += n
// 	}
// 	return sum
// }

// fmt.Println(Sum(1, 2, 3)) // → 6
