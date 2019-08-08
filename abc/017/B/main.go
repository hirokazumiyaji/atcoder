package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	i := len(x) - 1
	for i >= 0 {
		switch x[i] {
		case 'h':
			if i == 0 {
				fmt.Println("NO")
				return
			}
			if x[i-1] != 'c' {
				fmt.Println("NO")
				return
			}
			i -= 2
		case 'o', 'k', 'u':
			i--
		default:
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
