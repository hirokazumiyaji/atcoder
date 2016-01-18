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

	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	k, _ := strconv.Atoi(params[3])

	scanner.Scan()
	params = strings.Split(scanner.Text(), " ")
	s, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])

	result := (a * s) + (b * t)
	if s+t >= k {
		result -= (s + t) * c
	}
	fmt.Println(result)
}
