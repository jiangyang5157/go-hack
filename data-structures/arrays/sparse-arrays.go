package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	data := make(map[string]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		data[s]++
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < q; i++ {
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(data[s])
	}
}
