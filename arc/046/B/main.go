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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])

	for {
		if n-(a+b) <= (a + b) {
			break
		}
		n -= (a + b)
	}
	fmt.Println(n)

	fmt.Println("Takahashi")

	fmt.Println("Aoki")
}
