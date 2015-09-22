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
	result := 0
	for i := 1; i <= n; i++ {
		result += strings.Count(strconv.Itoa(i), "1")
	}
	fmt.Println(result)
}
