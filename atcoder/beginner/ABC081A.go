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
	s := scanner.Text()
	// fmt.Printf("s: %s\n", s)
	fmt.Println(strings.Count(s, "1"))

}
