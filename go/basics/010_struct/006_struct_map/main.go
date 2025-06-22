package main

import "fmt"

// 構造体　map
type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 20},
		2: {Name: "user2", Age: 30},
		3: {Name: "user3", Age: 40},
	}

	fmt.Println(m) // map[1:{user1 20} 2:{user2 30} 3:{user3 40}]

	m2 := map[User]string{
		{Name: "user1", Age: 20}: "Tokyo",
		{Name: "user2", Age: 25}: "Okayama",
	}

	fmt.Println(m2) // map[{user1 20}:Tokyo {user2 25}:Okayama]

	// makeで作成も可能
}
