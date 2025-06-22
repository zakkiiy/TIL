package main

import "fmt"

// コンストラクタ風の関数
// 「NewXxx()」という関数名を使って構造体の初期化関数を作る。
// ポインタで返すことで呼び出し側で更新可能

type User struct {
	Name string
	Age  int
}

// User構造体のインスタンスを生成して返す関数
// ポインタで返すことで、呼び出し側でフィールドを直接変更できるようにする（←Goの慣習）
// 値コピーを防げるので効率的（構造体が大きくなると特に重要）
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 20)
	fmt.Println(user1)  // &{user1 20}
	fmt.Println(*user1) // {user1 20}
}

/*
| 項目   | コンストラクタ関数（NewUser）        | 構造体メソッド（例：`func (u *User) Hello()`） |
| ---- | ------------------------- | ----------------------------------- |
| 役割   | インスタンスを**生成する**           | インスタンスを**操作する**                     |
| 書き方  | `func NewUser(...) *User` | `func (u *User) Xxx(...)`           |
| 所属   | 構造体の外に定義                  | 構造体のレシーバとして定義                       |
| 意味合い | 「ファクトリ関数」                 | 「振る舞い（メソッド）」                        |
*/

/*
	コンストラクタを使わない場合
	user := NewUser("user1", 20)
	手軽で短い
	シンプルな構造体にはこれで十分
	ただし、初期化ロジックがある場合には弱い
*/
/*
	コンストラクタを使う場合
	user := NewUser("user1", 20)
	「生成用の関数ですよ」という明確な意図を示せる
	内部で追加の処理（例：バリデーション、デフォルト値、ログなど）が書ける
	将来的に生成の仕組みが変わっても呼び出し側を修正しなくて済む
*/

/*
	func NewUser(name string, age int) *User {
		if name == "" {
			name = "anonymous" // ← デフォルトの名前に置き換え
		}
		return &User{Name: name, Age: age}
	}
	「名前が空だったら、自動的に 'anonymous' にする」
*/
