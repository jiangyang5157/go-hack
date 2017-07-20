package main

import (
	"fmt"
	"math"
)

func main() {
	solution([]int{4, 6, 2, 1, 5})
}

// h: The H
func solution(h []int) int {
	fmt.Println(h)

	// The N
	n := len(h)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Sorted H index
	sortedIndex := indexSort(h)
	fmt.Println(sortedIndex)

	// cache invalid result
	ret := 0
	for i := 0; i < len(sortedIndex); i++ {
		h[sortedIndex[i]] = int(math.Pow(2, float64(i)))
		fmt.Println(h[sortedIndex[i]])
		ret += h[sortedIndex[i]]
	}

	// calculate valid result
	for i := 0; i < len(sortedIndex); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < i; k++ {
				if k == j {
					continue
				}
				fmt.Println(sortedIndex, i, j, k, h, sortedIndex[i], sortedIndex[j], sortedIndex[k])

				// if sortedIndex[j] < sortedIndex[i] && sortedIndex[k] < sortedIndex[i] {
				// 	// left
				// 	if sortedIndex[k] > sortedIndex[j] {
				// 		fmt.Println("L")
				// 		ret -= h[sortedIndex[j]]
				// 	}
				// } else if sortedIndex[j] > sortedIndex[i] && sortedIndex[k] > sortedIndex[i] {
				// 	// right
				// 	if sortedIndex[k] > sortedIndex[j] {
				// 		fmt.Println("R")
				// 		ret -= h[sortedIndex[j]]
				// 	}
				// }
			}
		}
	}

	fmt.Println(ret)
	return ret
}

// Sort the arr, but returns the index
func indexSort(arr []int) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = i
	}
	_, ret = MergeIndexSort(arr, ret)
	return ret
}

func MergeIndexSort(arr, arrIndex []int) ([]int, []int) {
	if len(arr) <= 1 {
		return arr, arrIndex
	}

	middleIndex := len(arr) / 2
	leftPart, leftPartIndex := MergeIndexSort(arr[:middleIndex], arrIndex[:middleIndex])
	rightPart, rightPartIndex := MergeIndexSort(arr[middleIndex:], arrIndex[middleIndex:])
	return mergeIndex(leftPart, leftPartIndex, rightPart, rightPartIndex)
}

func mergeIndex(left, leftIndex, right, rightIndex []int) ([]int, []int) {
	leftLen := len(left)
	rightLen := len(right)
	ret := make([]int, 0, leftLen+rightLen)
	retIndex := make([]int, 0, leftLen+rightLen)
	for ; leftLen > 0 || rightLen > 0; leftLen, rightLen = len(left), len(right) {
		if leftLen == 0 {
			return append(ret, right...), append(retIndex, rightIndex...)
		}
		if rightLen == 0 {
			return append(ret, left...), append(retIndex, leftIndex...)
		}

		if right[0] > left[0] {
			ret = append(ret, right[0])
			retIndex = append(retIndex, rightIndex[0])
			right = right[1:]
			rightIndex = rightIndex[1:]
		} else {
			ret = append(ret, left[0])
			retIndex = append(retIndex, leftIndex[0])
			left = left[1:]
			leftIndex = leftIndex[1:]
		}
	}
	return ret, retIndex
}
