package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	l, _ := strconv.Atoi(params[0])
	h, _ := strconv.Atoi(params[1])

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		if l <= a && a <= h {
			fmt.Println(0)
		} else if a > h {
			fmt.Println(-1)
		} else {
			fmt.Println(l - a)
		}
	}
}
