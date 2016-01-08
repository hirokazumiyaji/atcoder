package main

import (
	"bufio"
	"fmt"
	"math"
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

	if n >= 12 {
		n -= 12
	}
	diff := int32(math.Abs(float64(n*30) - float64(m*6)))
	if diff > 180 {
		fmt.Println(diff - 180)
	} else {
		fmt.Println(diff)
	}
}
