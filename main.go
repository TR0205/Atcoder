package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	res := "Yes"
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = "No"
		}
	}
	fmt.Println(res)
}
