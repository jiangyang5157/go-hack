//N voracious fish are moving along a river. Calculate how many fish are alive.
//https://codility.com/programmers/lessons/7-stacks_and_queues/fish/
package main

import "fmt"

func main() {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}
	Num := Solution(A, B)
	fmt.Printf("A=%v, B=%v\nNum=%v\n", A, B, Num)
}

func Solution(A []int, B []int) int {
	return 0
}
