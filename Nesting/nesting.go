//Determine whether a given string of parentheses is properly nested.
//https://codility.com/programmers/lessons/7-stacks_and_queues/nesting/
//It is a special case of the Brackets problem
package main

import "fmt"

func main() {
	S := "(()(())())"
	// S := "())"
	IsNested := Solution(S)
	fmt.Printf("S=%v\nIsNested=%v\n", S, IsNested)
}

func Solution(S string) int {
	if len(S)%2 != 0 {
		return 0
	}

	open := 0
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case '(':
			open++
		case ')':
			open--
		}
		if open < 0 {
			return 0
		}
	}

	if open == 0 {
		return 1
	}
	return 0
}
