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
	n, _ := strconv.Atoi(scanner.Text())

	if n == 100 {
		fmt.Println("Perfect")
	} else if n >= 90 {
		fmt.Println("Great")
	} else if n >= 60 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
