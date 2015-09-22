package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for i := 0; i < 12; i++ {
		scanner.Scan()
		if strings.Contains(scanner.Text(), "r") {
			count += 1
		}
	}
	fmt.Println(count)
}
