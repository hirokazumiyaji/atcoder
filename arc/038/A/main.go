package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	aList := make([]int, n)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		aList[i], _ = strconv.Atoi(v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(aList)))

	count := 0
	for i := 0; i < n; i += 2 {
		count += aList[i]
	}

	fmt.Println(count)
}
