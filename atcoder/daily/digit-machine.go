package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}

	now := 0
	for i := 0; i < 3; i++ {
		now = a[now]
	}
	fmt.Println(now)
}
