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
	line := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	r, _ := strconv.Atoi(line[1])
	fmt.Println(n, r)
	scanner.Scan()
	s := scanner.Text()
	count := len(s)
	drawCount := r - 1
	for i := 0; i < count; i++ {
		if strings.Count(s, "o") == count {
			fmt.Println(i)
		}
	}
}
