//Find the index S such that the leaders of the sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N - 1] are the same.
//https://codility.com/programmers/lessons/8-leader/equi_leader/
package main

import "fmt"

func main() {
	A := []int{4, 3, 4, 4, 4, 2}
	result := Solution(A)
	fmt.Printf("A=%v\nresult=%v\n", A, result)
}

func Solution(A []int) int {
	return -1
}
