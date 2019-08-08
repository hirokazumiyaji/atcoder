package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	plist := make([]int, 0, len(parts))
	for _, v := range parts {
		i, _ := strconv.Atoi(v)
		plist = append(plist, i)
	}
}
