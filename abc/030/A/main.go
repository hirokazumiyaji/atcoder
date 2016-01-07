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
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])

	wp_t := float32(b) / float32(a)
	wp_a := float32(d) / float32(c)
	if wp_t == wp_a {
		fmt.Println("DRAW")
	} else if wp_t > wp_a {
		fmt.Println("TAKAHASHI")
	} else {
		fmt.Println("AOKI")
	}
}
