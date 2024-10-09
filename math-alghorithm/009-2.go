package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Scanf("%d %d", &n, &s)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	// 0からn枚目までのスライス
	dp := make([][]bool, n+1)

	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if a[i] > j {
				dp[i][j] = dp[i-1][j]
			}
			if a[i] <= j {
				if dp[i-1][j] || dp[i-1][j-a[i]] {
					dp[i][j] = true
				} else {
					dp[i][j] = false

				}
			}
		}
	}

	res := ""
	if dp[n][s] {
		res = "Yes"
	} else {
		res = "No"

	}
	fmt.Println(res)
}
