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
	result := ""
	var char rune
	count := 0
	for i, r := range s {
		if r == char {
			count++
			continue
		}
		if i != 0 {
			result += string([]rune{char}) + strconv.Itoa(count)
		}
		char = r
		count = 1
	}
	result += string([]rune{char}) + strconv.Itoa(count)
	fmt.Println(result)
}
