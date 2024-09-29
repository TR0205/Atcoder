package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}

	fmt.Println(sum % 100)
}
