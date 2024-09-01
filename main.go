package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	r := make([]bool, n)

	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)

		// fmt.Printf("%d, %d, %d", t, x, y)

		// 2の場合、y=-x+2, y=-x
		// 3の場合、y=-x+3, y=-x+1
		// 6の場合、y=-x+6, y=-x+4, y=-x+2, y=-x

		for t >= 0 {
			if y == -x+t {
				r[i] = true
				break
			}
			t -= 2
		}
		if t == -1 {
			r[i] = false
		}
	}

	// fmt.Println(r)

	var j bool = true
	for _, v := range r {
		if !v {
			j = false
		}
	}

	if j {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")

	}
}
