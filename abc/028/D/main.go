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
	n, _ := strconv.Atoi(params[0])
	k, _ := strconv.Atoi(params[1])

	if k == 1 {
		fmt.Println(
			((float64(1) / float64(n)) * (float64(1) / float64(n)) * (float64(1) / float64(n))) +
				((float64(1) / float64(n)) * (float64(1) / float64(n)) * (float64(n-1) / float64(n)) * 3),
		)
	} else {
		fmt.Println(
			((float64(1) / float64(n)) * (float64(1) / float64(n)) * (float64(1) / float64(n))) +
				((float64(1) / float64(n)) * (float64(k) / float64(n)) * (float64((n - k)) / float64(n)) * 3) +
				((float64(1) / float64(n)) * (float64(1) / float64(n)) * (float64(n-1) / float64(n)) * 3),
		)
	}
}
