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
	n, _ := strconv.Atoi(scanner.Text())
	results := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		_, _ = strconv.Atoi(scanner.Text()) // Replace x with _
		sum := 0

		scanner.Scan()
		nums := strings.Fields(scanner.Text())
		for _, num := range nums {
			y, _ := strconv.Atoi(num)
			if y >= 0 {
				sum += y * y
			}
		}
		results[i] = sum
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
