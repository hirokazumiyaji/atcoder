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

	scanner.Scan()
	params = strings.Split(scanner.Text(), " ")
	x, _ := strconv.Atoi(params[0])
	y, _ := strconv.Atoi(params[1])

	aList := make([]int, n)
	bList := make([]int, m)

	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		aList[i], _ = strconv.Atoi(v)
	}

	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		bList[i], _ = strconv.Atoi(v)
	}

	hour := 0
	current := 0
	count := 0
	aCurrent := 0
	bCurrent := 0
	for {
		if current == 0 {
			count += 1
			if aList[n-1] < hour {
				break
			}

			hit := false
			for i := aCurrent; i < n; i++ {
				if aList[i] >= hour {
					aCurrent = i
					hour = aList[i] + x
					hit = true
					break
				}
			}
			if hit == false {
				break
			}

			current = 1
		} else {
			if bList[m-1] < hour {
				break
			}

			hit := false
			for i := bCurrent; i < m; i++ {
				if bList[i] >= hour {
					bCurrent = i
					hour = bList[i] + y
					hit = true
					break
				}
			}
			if hit == false {
				break
			}

			current = 0
		}
	}

	fmt.Println(count - 1)
}
