package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		s = s[:l-1] + reverse(s[l-1:r]) + s[r:]
	}
	fmt.Println(s)
}
