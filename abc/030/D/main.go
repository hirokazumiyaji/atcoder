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
	params := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(params[0])
	a, _ := strconv.Atoi(params[1])

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	bList := make([]int, n)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		bList[i], _ = strconv.Atoi(v)
	}
}
