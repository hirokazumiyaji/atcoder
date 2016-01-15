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
	s := scanner.Text()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	if n < 5 {
		fmt.Println(fmt.Sprintf("%c%c", s[0], s[n-1]))
	} else if n < 10 {
		fmt.Println(fmt.Sprintf("%c%c", s[1], s[n-5-1]))
	} else if n < 15 {
		fmt.Println(fmt.Sprintf("%c%c", s[2], s[n-10-1]))
	} else if n < 20 {
		fmt.Println(fmt.Sprintf("%c%c", s[3], s[n-15-1]))
	} else {
		fmt.Println(fmt.Sprintf("%c%c", s[4], s[n-20-1]))
	}
}
