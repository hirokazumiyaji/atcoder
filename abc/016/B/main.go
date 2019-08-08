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
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])
	if a+b == c && a-b == c {
		fmt.Println("?")
	} else if a+b == c {
		fmt.Println("+")
	} else if a-b == c {
		fmt.Println("-")
	} else {
		fmt.Println("!")
	}
}
