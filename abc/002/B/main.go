package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	w := scanner.Text()
	s := make([]rune, 0, len(w))
	for _, r := range w {
		switch r {
		case 'a', 'i', 'u', 'e', 'o':
		default:
			s = append(s, r)
		}
	}
	fmt.Println(string(s))
}
