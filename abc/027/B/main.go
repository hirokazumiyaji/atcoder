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

	aList := make([]int, n)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		aList[i], _ = strconv.Atoi(v)
	}
}
