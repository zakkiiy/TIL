package main

import "fmt"

// 構造体　スライス
type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}

	// make
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)
	for _, u := range users2 {
		fmt.Println(*u)
	}
}

/*
	構造体（User）のポインタをスライス（[]*User）で管理することで、複数のユーザー情報を効率的に扱える
	よくある「JSON配列で複数のオブジェクトを返す」パターンに対応
	makeやappendを使って柔軟にデータを追加・操作可能
*/
