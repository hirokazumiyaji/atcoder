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
	lines := make(map[int]int)
	for _, v := range strings.Split(scanner.Text(), " ") {
		n, _ := strconv.Atoi(v)
		lines[n] += 1
	}
	result := 0
	for k := range lines {
		if lines[k] == 3 {
			result = k
			break
		} else if lines[k] == 1 {
			result = k
			break
		}
	}
	fmt.Println(result)
}
