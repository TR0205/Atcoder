package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	charSlice := strings.Split(input, "")

	count1 := 0
	count2 := 0
	count3 := 0
	for i := 0; i < len(charSlice); i++ {
		if charSlice[i] == "1" {
			count1++
		}
		if charSlice[i] == "2" {
			count2++
		}
		if charSlice[i] == "3" {
			count3++
		}
	}

	if count1 == 1 && count2 == 2 && count3 == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")

	}
}
