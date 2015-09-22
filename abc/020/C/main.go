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
	numbers := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(numbers[0])
	w, _ := strconv.Atoi(numbers[1])
	t, _ := strconv.Atoi(numbers[2])
	fmt.Println(h)
	fmt.Println(w)
	fmt.Println(t)
}
