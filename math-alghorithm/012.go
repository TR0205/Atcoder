package main

import "fmt"

// nが素数かどうかを判定する
// n= 53の場合「2から⌊√53⌋までの数字を割ってあまりが0になるものがなければ素数」と判定できる
// ⌊√53⌋は7.***以下で最大の整数を表しており、この場合は7
// 2,3,4,5,6,7まで判定する
func main() {
	var n int
	fmt.Scanf("%d", &n)

	res := "Yes"
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = "No"
		}
	}
	fmt.Println(res)
}
