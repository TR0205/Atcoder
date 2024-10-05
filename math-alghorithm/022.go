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

	// aの各要素がそれぞれ何個あるかをカウントするスライスを定義
	cnt := make([]int, 100000)

	for i := 0; i < n; i++ {
		// 例：a[i] = 1の場合c[1] += 1とする
		cnt[a[i]] += 1
	}

	res := 0
	// 合計100,000となる組み合わせを一つずつ確認していく
	// 1が2件、99999が3件ある場合、2*3=6をresに加える
	for i := 1; i <= 50000; i++ {
		if cnt[i] > 0 && cnt[100000-i] > 0 {
			// 50,000+50,000の場合
			if i == 50000 {
				// nC2の組み合わせとなるため単純な積ではなくなる
				// 例：50,000が4件ある場合、4C2の組み合わせとなり4*3/2!のようになる
				res += (cnt[i] * (cnt[i] - 1)) / 2
				continue
			}
			res += cnt[i] * cnt[100000-i]
		}
	}

	fmt.Println(res)
}
