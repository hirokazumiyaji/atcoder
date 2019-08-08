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
	total := 0
	for i := 0; i < 3; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		s, _ := strconv.Atoi(parts[0])
		e, _ := strconv.Atoi(parts[1])
		total += s * e / 10
	}
	fmt.Println(total)
}
