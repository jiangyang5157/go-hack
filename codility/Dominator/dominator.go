//Find an index of an array such that its value occurs at more than half of indices in the array.
//https://codility.com/programmers/lessons/6-sorting/distinct/
package main

import "fmt"

func main() {
	A := []int{3, 4, 3, 2, 3, -1, 3, 3}
	Index := Solution(A)
	fmt.Printf("A=%v\nIndex=%v\n", A, Index)
}

func Solution(A []int) int {
	return -1
}
