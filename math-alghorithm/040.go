package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	strA := strings.Fields(input)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			a[i] = 0
			continue
		}
		a[i], _ = strconv.Atoi(strA[i-1])
	}

	// 累積和
	cum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			cum[0] = a[0]
			continue
		}
		cum[i] = cum[i-1] + a[i]
	}

	// mの入力処理
	inputM, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(inputM))

	b := make([]int, m)
	res := 0
	for i := 0; i < m; i++ {
		inputB, _ := reader.ReadString('\n')
		bin, _ := strconv.Atoi(strings.TrimSpace(inputB))

		if i == 0 {
			b[0] = bin
			continue
		}

		b[i] = bin
		res += abs(cum[(b[i]-1)] - cum[(b[i-1]-1)])

	}

	fmt.Println(res)

}
