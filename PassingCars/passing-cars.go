//Count the number of passing cars on the road.
//https://codility.com/programmers/lessons/5-prefix_sums/passing_cars/
package main

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}
	Counts := Solution(A)
	fmt.Printf("A=%v\nCounts=%v\n", A, Counts)
}

func Solution(A []int) int {
	ret := 0
	tmp := 0
	for _, v := range A {
		if v == 0 {
			tmp++
		} else if v == 1 {
			ret += tmp
		}
		if ret > 1000000000 {
			return -1
		}
	}
	return ret
}
