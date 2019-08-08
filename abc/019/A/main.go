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
	parts := strings.Split(scanner.Text(), " ")
	list := make([]int, 0, 3)
	for _, e := range parts {
		i, _ := strconv.Atoi(e)
		list = append(list, i)
	}
	sort.Ints(list)
	fmt.Println(list[1])
}
