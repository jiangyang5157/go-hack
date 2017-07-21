//Find the earliest time when a frog can jump to the other side of a river.
//https://codility.com/programmers/lessons/4-counting_elements/frog_river_one/
package main

import "fmt"

func main() {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	EarliestTime := Solution(X, A)
	fmt.Printf("X=%v, A=%v\nEarliestTime=%v\n", X, A, EarliestTime)
}

func Solution(X int, A []int) int {
	ret := -1
	aMap := make(map[int]int)
	for i, v := range A {
		if v > X {
			continue
		}
		aMap[v]++
		if len(aMap) == X {
			ret = i
			break
		}
	}
	return ret
}
