package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	cycle, result := 1, 0
	for i := 0; i < 9; i++ {
		num := (n / cycle) % 10
		cycle *= 10
		result += (n / cycle) * (cycle / 10)
		switch {
		case num == 1:
			result += (n % (cycle / 10)) + 1
		case num > 1:
			result += cycle / 10
		}
	}
	fmt.Println(result)
}
