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
	n, _ := strconv.Atoi(scanner.Text())
	tc := 0
	ac := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		for _, s := range scanner.Text() {
			if s == 'R' {
				tc += 1
			} else if s == 'B' {
				ac += 1
			}
		}
	}

	if tc > ac {
		fmt.Println("TAKAHASHI")
	} else if ac > tc {
		fmt.Println("AOKI")
	} else {
		fmt.Println("DRAW")
	}
}
