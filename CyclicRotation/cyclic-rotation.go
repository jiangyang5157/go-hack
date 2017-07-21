//Rotate an array to the right by a given number of steps.
//https://codility.com/programmers/lessons/2-arrays/cyclic_rotation/
package main

import "fmt"

func main() {
	input1 := []int{3, 8, 9, 7, 6}
	input2 := 3
	output := Solution(input1, input2)
	fmt.Printf("A=%v, K=%v\nA'=%v\n", input1, input2, output)
}

func Solution(A []int, K int) []int {
	if len(A) > 1 {
		for K > 0 {
			var x int
			//pop
			x, A = A[len(A)-1], A[:len(A)-1]
			//push back
			A = append([]int{x}, A...)
			K--
		}
	}
	return A
}
