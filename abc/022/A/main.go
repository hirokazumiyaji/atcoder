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
	line := scanner.Text()
	elements := strings.Split(line, " ")
	n, _ := strconv.Atoi(elements[0])
	s, _ := strconv.Atoi(elements[1])
	t, _ := strconv.Atoi(elements[2])
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	result := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		w += a
		if s <= w && w <= t {
			result += 1
		}
	}
	fmt.Println(result)
}
