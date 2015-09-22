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
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := scanner.Text()
	xy := make([]int, 2)
	for i := 0; i < k; i++ {
		if s[i] == 'L' {
			xy[0] -= (i + 1)
		} else if s[i] == 'R' {
			xy[0] += (i + 1)
		} else if s[i] == 'U' {
			xy[1] += (i + 1)
		} else if s[i] == 'D' {
			xy[1] -= (i + 1)
		}
	}
	fmt.Println(xy[0], xy[1])
}
