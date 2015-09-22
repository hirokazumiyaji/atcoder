package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mem = make(map[[2]int]int, 100)

func gcd(x, y int) int {
	n := [2]int{x, y}

	if y == 0 {
		return x
	}
	if m, ok := mem[n]; ok {
		return m
	}
	m := gcd(y, x%y)
	mem[n] = m
	return m
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(numbers[0])
	k, _ := strconv.Atoi(numbers[1])
	sum := 0
	for i := 1; i <= n; i++ {
		sum += lcm(i, k)
	}
	fmt.Println(sum % 1000000007)
}
