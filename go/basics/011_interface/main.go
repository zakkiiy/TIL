package main

import "fmt"

// これは「ToString というメソッドを持ってる型しか、この interface に該当しないよ」という “型の契約”
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

// 両方の構造体に、ToString() メソッドを実装
//これで：
// *Person は Stringfy インターフェースに準拠している
// *Car も Stringfy インターフェースに準拠している

func (p *Person) ToString() string {
	return fmt.Sprintln("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintln("Name=%v, Age=%v", c.Number, c.Model)
}

// この interface 型のスライスを作る
// Stringfy という型のスライスを作っている
// 中身は *Person, *Car のインスタンス（オブジェクト）だけど、interface として見ている
func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 20},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	// なので、forループで全体を回してもどちらの型でも .ToString() が呼ばれる
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// Goでは構造体が明示的に「implements Stringfy」と書かなくてもOK
}

/*
型が違っても、同じメソッド（ToString）を使いたい」から！
たとえば []Person と []Car は、型が違うから一つのスライスに入れられない
→ でも「ToString() を持っているもの」という共通点でひとまとめにしたい
→ そのために interface を使う！

Person と Car という構造体に、同じ名前のメソッド（ToString）を定義
その共通点を元に interface 型のスライス []Stringfy を作成
vs の中身はそれぞれ別の構造体だが、Stringfy として統一的に扱える！
*/
