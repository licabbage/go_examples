package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用$cat ./dat1 | go run src/line_filters.go 来运行。
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
