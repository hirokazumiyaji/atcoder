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
	i, _ := strconv.Atoi(scanner.Text())
	if i == 1 {
		fmt.Println("ABC")
	} else {
		fmt.Println("chokudai")
	}
}
