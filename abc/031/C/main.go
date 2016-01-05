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

	n, _ := strconv.Atoi(scanner.Text())
	as := make([]int, n)
	for i, v := range strings.Split(scanner.Text(), " ") {
		as[i], _ = strconv.Atoi(v)
	}
	fmt.Println(as)
}
