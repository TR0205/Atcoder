package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var min int
	var max int

	if scanner.Scan() {
		line := scanner.Text()
		// 空白(スペースやタブ)を区切り文字として文字列を分割
		fields := strings.Fields(line)
		num, _ = strconv.Atoi(fields[0])
		min, _ = strconv.Atoi(fields[1])
		max, _ = strconv.Atoi(fields[2])
	}

	total := 0
	for i := 0; i <= num; i++ {
		sum := sumOfDigits(i)
		if sum >= min && sum <= max {
			total += i
		}
	}
	fmt.Println(total)
}
