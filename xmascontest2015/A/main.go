package main

import (
	"bufio"
	"fmt"
	"math"
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
	x, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])
	a, _ := strconv.Atoi(params[2])
	b, _ := strconv.Atoi(params[3])
	c, _ := strconv.Atoi(params[4])

	s := 0
	for i := 1; i <= n; i++ {
		s += x
		for j := 1; j <= t; j++ {
			x = int(math.Mod(float64(a*x+b), float64(c)))
		}
	}

	fmt.Println(s)
}
