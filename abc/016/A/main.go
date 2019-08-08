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
	parts := strings.Split(scanner.Text(), " ")
	m, _ := strconv.Atoi(parts[0])
	d, _ := strconv.Atoi(parts[1])
	if m%d == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
