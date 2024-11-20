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

// スライスを空白区切りで出力。Printlnだと改行されてしまう
func outputSlice() {
	for _, v := range a {
		fmt.Print(v, " ")
	}
}

func abouRune() {
	// 「あ」はUnicodeでU+3042と表す
	// 3042部分がcode point
	// Goが文字を扱う場合はUTF-8という方式で文字列を置き換えている(符号化方式)
	// UTF-8で「あ」は「e3 81 82」のように16進数表記の3バイトで表現される
	// そのためs[i]などで出力すると以下のようになる
	s := "あ"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i]) // %xはフォーマット指定子
	}
	// 出力結果：
	// e3
	// 81
	// 82

	// 単純に出力すると10進数表記のバイト値が出力される
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	// 出力結果：
	// 227
	// 129
	// 130

	// runeは「e3 81 82」のようにbyte単位で扱うのではなく、code point単位で「あ」として扱う
	// 以下の例だと1文字として扱うため、1となる
	fmt.Println(len([]rune(s)))
	// 出力結果：
	// 1
}
