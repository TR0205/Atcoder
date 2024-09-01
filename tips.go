package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 文字列を整数の配列へ変換
func createArr() {
	scanner := bufio.NewScanner(os.Stdin)

	var nums []int
	if scanner.Scan() {
		line := scanner.Text()
		// 空白(スペースやタブ)を区切り文字として文字列を分割
		fields := strings.Fields(line)
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}
	}
}

// スライスの生成(例：5 6 8 10のようなinput)
func createSlice() {
	n := 4
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan((&a[i]))
	}

	fmt.Println(a)
}
