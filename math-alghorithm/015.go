package main

import "fmt"

// ユークリッドの互助法を使用した最大公約数の計算
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// 大きい数/小さい数の余りを大きい数と入れ替え、どちらかが0になるまで繰り返す
	for a >= 1 && b >= 1 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	// 0ではない法の数を返す
	if a >= 1 {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
