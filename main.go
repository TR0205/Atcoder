package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan((&a[i]))
	}

	counts := make([]int, n)
	for i, v := range a {
		counts[i] = 0
		// 条件を満たしている間ループし続ける
		for v%2 == 0 {
			v /= 2
			counts[i]++
		}
	}

	// AtCoderで使われているGoのバージョン1.20.6だと存在しない？
	// fmt.Println(slices.Max(counts))
	// fmt.Println(counts)

	// 以下で代用
	min := counts[0]
	for _, v := range counts {
		if min > v {
			min = v
		}
	}
	fmt.Println(min)

	// scanner := bufio.NewScanner(os.Stdin)
	// // scanner.Scan()
	// // s := scanner.Text()

	// // fmt.Println(strings.Count(s, "1"))

	// var count int
	// if scanner.Scan() {
	// 	line := scanner.Text()
	// 	// 空白(スペースやタブ)を区切り文字として文字列を分割
	// 	fields := strings.Fields(line)
	// 	for _, field := range fields {
	// 		for field % 2 != 0 {
	// 			count++
	// 		}
	// 	}
	// }

	// num := nums[0]
	// total := nums[1]

	// for i := 0; i < num; i++ {
	// 	for j := 0; j < num; j++ {
	// 		k := num - i - j
	// 		if 10000*i+5000*j+1000*k == total && k >= 0 {
	// 			fmt.Printf("%d %d %d\n", i, j, k)
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println("-1 -1 -1\n")

	// i := []int{5, 3, 2, 8, 7}
	// // s := []string{"d", "a", "f"}

	// fmt.Println(i)
	// // 昇順でソート
	// sort.Ints(i)
	// fmt.Println(i)

	// // 降順でソート
	// sort.Sort(sort.Reverse(sort.IntSlice(i)))
	// fmt.Println(i)

}
