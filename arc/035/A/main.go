package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	success := true
	for scanner.Scan() {
		s = scanner.Text()
	}

	length := len(s)

	for i := 0; i < length; i++ {
		if s[i] == s[length-i-1] {
			continue
		}
		if s[i] == '*' || s[length-i-1] == '*' {
			continue
		}
		success = false
		break
	}

	if success {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
