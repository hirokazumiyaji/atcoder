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
	input := scanner.Text()
	numbers := strings.Split(input, " ")
	num, _ := strconv.Atoi(numbers[0] + numbers[1])
	fmt.Println(num * 2)
}
