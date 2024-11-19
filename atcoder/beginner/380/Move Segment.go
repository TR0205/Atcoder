package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	charSlice := strings.Split(input, "")

	var a []string
	c := "0"
	for i, v := range charSlice {
		if i == 0 {
			c = v
			a = append(a, v)
			continue
		}

		if c != v {
			a = append(a, v)
			c = v
			continue
		}

		a[len(a)-1] += v
	}

	idx := 0
	if charSlice[0] == "0" {
		idx = k*2 - 1
	} else {
		idx = k*2 - 2
	}

	a[idx-1], a[idx] = a[idx], a[idx-1]

	fmt.Println(strings.Join(a, ""))

}
