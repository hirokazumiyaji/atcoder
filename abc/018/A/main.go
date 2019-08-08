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
	data := make([]int, 0, 3)
	sdata := make([]int, 0, 3)
	m := make(map[int]int)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		data = append(data, n)
		sdata = append(sdata, n)
		m[n] = 0
	}
	sort.Ints(sdata)
	i := 3
	for _, n := range sdata {
		m[n] = i
		i--
	}
	for _, n := range data {
		fmt.Println(m[n])
	}
}
