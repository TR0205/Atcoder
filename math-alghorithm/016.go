package main

import "fmt"

// ユークリッドの互助法を使用した最大公約数の計算
func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		for res >= 1 && a[i] >= 1 {
			if res > a[i] {
				res = res % a[i]
			} else {
				a[i] = a[i] % res
			}
		}
		if a[i] > 0 {
			res = a[i]
		}
	}
	fmt.Println(res)
}
