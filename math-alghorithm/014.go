package main

import "fmt"

// 素因数分解のコード
// n=53の場合、⌊√53⌋までの値を割っていき、余りが0の場合はスライスに値を追加する
// 2,3,4,5,6,7までループで確認する
func main() {
	var n int
	fmt.Scanf("%d", &n)

	var res []int
	for i := 2; i*i <= n; i++ {
		// 余りが0の場合はnをiで割り、nへ代入
		for n%i == 0 {
			n /= i
			res = append(res, i)
		}
	}
	// 最後に残ったnが1より上の場合、その値を出力結果へ加える
	if n > 1 {
		res = append(res, n)
	}

	// ループで半角区切りで出力
	for _, v := range res {
		fmt.Print(v, " ")
	}
}
