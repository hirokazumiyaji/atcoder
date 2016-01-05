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
	params := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(params[0])
	d, _ := strconv.Atoi(params[1])
	result1 := (a + 1) * d
	result2 := a * (d + 1)
	if result1 > result2 {
		fmt.Println(result1)
	} else {
		fmt.Println(result2)
	}
}
