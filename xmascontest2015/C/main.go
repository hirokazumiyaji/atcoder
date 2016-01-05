package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params = strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(params[0])
	w, _ := strconv.Atoi(params[1])
}
