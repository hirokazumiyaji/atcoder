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
	elements := strings.Split(scanner.Text(), " ")
	x, _ := strconv.ParseInt(elements[0], 10, 64)
	y, _ := strconv.ParseInt(elements[1], 10, 64)
	if x > y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}
