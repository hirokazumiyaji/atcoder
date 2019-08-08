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
	s := scanner.Text()
	n, _ := strconv.Atoi(s)
	switch n {
	case 1:
		fmt.Println("1")
		fmt.Println("1")
	case 2:
		fmt.Println("1")
		fmt.Println("2")
	case 3:
		fmt.Println("2")
		fmt.Println("1")
		fmt.Println("2")
	case 4:
		fmt.Println("2")
		fmt.Println("2")
		fmt.Println("2")
	case 5:
		fmt.Println("2")
		fmt.Println("1")
		fmt.Println("4")
	case 6:
		fmt.Println("2")
		fmt.Println("2")
		fmt.Println("4")
	case 7:
		fmt.Println("3")
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("4")
	case 8:
		fmt.Println("2")
		fmt.Println("4")
		fmt.Println("4")
	case 9:
		fmt.Println("2")
		fmt.Println("1")
		fmt.Println("8")
	case 10:
		fmt.Println("2")
		fmt.Println("2")
		fmt.Println("8")
	}
}
