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
	ab := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(ab[0])
	b, _ := strconv.Atoi(ab[1])
	ans := -3000

	t := (a%100 + 900) - b
	if ans < t {
		ans = t
	}
	t = (a - (a / 10 % 10 * 10) + 90) - b
	if ans < t {
		ans = t
	}
	t = (a - (a % 10) + 9) - b
	if ans < t {
		ans = t
	}

	t = a - (b%100 + 100)
	if ans < t {
		ans = t
	}
	t = a - (b - (b / 10 % 10 * 10))
	if ans < t {
		ans = t
	}
	t = a - (b - (b % 10))
	if ans < t {
		ans = t
	}

	fmt.Println(ans)
}
