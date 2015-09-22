package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	result := 0
	var numbers []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		if sort.SearchInts(numbers, a) != len(numbers)-1 {
			result += 1
		} else {
			numbers = append(numbers, a)
			sort.Ints(numbers)
		}
	}
	fmt.Println(result)
}
