package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()

	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0

	for _, v := range s {
		switch v {
		case 'A':
			a += 1
		case 'B':
			b += 1
		case 'C':
			c += 1
		case 'D':
			d += 1
		case 'E':
			e += 1
		case 'F':
			f += 1
		}
	}
	fmt.Println(fmt.Sprintf("%d %d %d %d %d %d", a, b, c, d, e, f))
}
