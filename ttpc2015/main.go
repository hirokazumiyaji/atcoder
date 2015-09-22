package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputs := strings.SplitN(input, " ", 3)
	N, _ := strconv.Atoi(inputs[0])
	B, _ := strconv.Atoi(inputs[1])
	C, _ := strconv.Atoi(inputs[2])
	A := make([]int, 0, N)

	scanner.Scan()
	input = scanner.Text()
	for _, s := range strings.SplitN(input, " ", N) {
		i, _ := strconv.Atoi(s)
		A = append(A, i)
	}

}
