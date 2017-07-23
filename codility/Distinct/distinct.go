//Compute number of distinct values in an array.
//https://codility.com/programmers/lessons/6-sorting/distinct/
package main

import "fmt"

func main() {
	A := []int{2, 1, 1, 2, 3, 1}
	Num := Solution(A)
	fmt.Printf("A=%v\nNum=%v\n", A, Num)
}

func Solution(A []int) int {
	aMap := make(map[int]int)
	for _, v := range A {
		aMap[v]++
	}
	return len(aMap)
}
