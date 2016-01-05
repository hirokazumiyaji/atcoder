package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n, _ := strconv.Atoi(scanner.Text())

	as := make([]int, n)
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		as[i], _ = strconv.Atoi(params[i])
	}
}
