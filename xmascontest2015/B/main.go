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
	m, _ := strconv.Atoi(params[1])
	ps := make([]int, m)
	qs := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		params = strings.Split(scanner.Text(), " ")
		ps[i], _ = strconv.Atoi(params[0])
		qs[i], _ = strconv.Atoi(params[1])
	}
	fmt.Println(n)
	fmt.Println(ps)
	fmt.Println(qs)
}
