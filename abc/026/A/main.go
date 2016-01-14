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
	a, _ := strconv.Atoi(scanner.Text())
	half := a / 2
	fmt.Println(half * (a - half))
}
