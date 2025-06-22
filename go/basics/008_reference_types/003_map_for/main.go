package main

import "fmt"

func main() {
	m := map[string]int{
		"apple": 500,
		"lemon": 700,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
