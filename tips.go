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

// 二次元スライスの生成
func generate2DimensionSlice() {
	n := 3
	s := 11
	// 3*11のスライスを生成
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, s)
	}
}

// ランダムな文字数の文字列をスライスへ格納
func stringSlice() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	charSlice := strings.Split(input, "")

	fmt.Println(charSlice)
}
