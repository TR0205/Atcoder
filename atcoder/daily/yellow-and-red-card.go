package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	count := make([]int, n)

	for i := 0; i < q; i++ {
		var k, x int
		fmt.Scanf("%d %d", &k, &x)

		if k == 3 && count[x-1] > 1 {
			fmt.Println("Yes")
			continue
		}
		if k == 3 && count[x-1] <= 1 {
			fmt.Println("No")
			continue
		}
		if k == 1 {
			count[x-1] += 1
			continue
		}
		if k == 2 {
			count[x-1] += 2
			continue
		}
	}

}
