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
	params := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])
	e, _ := strconv.Atoi(params[4])

	sumMap := make(map[int]bool)
	sumMap[a+b+c] = true
	sumMap[a+b+d] = true
	sumMap[a+b+e] = true
	sumMap[a+c+d] = true
	sumMap[a+c+e] = true
	sumMap[a+d+e] = true
	sumMap[b+c+d] = true
	sumMap[b+c+e] = true
	sumMap[b+d+e] = true
	sumMap[c+d+e] = true

	sumList := make([]int, 10)
	for k := range sumMap {
		sumList = append(sumList, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sumList)))
	fmt.Println(sumList[2])
}
