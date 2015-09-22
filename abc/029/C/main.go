package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var chars []string = []string{"a", "b", "c"}

func permurations(prefix string, index, length int) {
	if index >= length {
		fmt.Println(prefix)
	} else {
		index += 1
		for _, s := range chars {
			permurations(prefix+s, index, length)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	permurations("", 0, n)
}
