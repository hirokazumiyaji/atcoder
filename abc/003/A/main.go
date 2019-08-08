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
	result := int64(0)
	for i := 0; i < n; i++ {
		result += int64((i + 1) * 10000 * (1 / (n + 1)))
	}
	fmt.Println(result)
}
