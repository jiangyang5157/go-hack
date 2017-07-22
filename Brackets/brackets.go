//Determine whether a given string of parentheses is properly nested.
//https://codility.com/programmers/lessons/7-stacks_and_queues/brackets/
package main

import "fmt"

func main() {
	S := "{[()()]}"
	// S := "([)()]"
	IsNested := Solution(S)
	fmt.Printf("S=%v\nIsNested=%v\n", S, IsNested)
}

func Solution(S string) int {
	s := []byte{}
	if len(S)%2 != 0 {
		return 0
	}

	for i := 0; i < len(S); i++ {
		switch S[i] {
		case ')':
			s = process(s, '(')
		case ']':
			s = process(s, '[')
		case '}':
			s = process(s, '{')
		default:
			s = Push(s, S[i])
		}
	}
	if len(s) > 0 {
		return 0
	}

	return 1
}

func process(s []byte, b byte) []byte {
	if !IsEmpty(s) && Peek(s) == b {
		_, s = Pop(s)
	}
	return s
}

func Length(s []byte) int {
	return len(s)
}

func IsEmpty(s []byte) bool {
	return len(s) == 0
}

func Peek(s []byte) byte {
	return s[len(s)-1]
}

func Pop(s []byte) (byte, []byte) {
	return s[len(s)-1], s[:len(s)-1]
}

func PopBack(s []byte) (byte, []byte) {
	return s[0], s[1:]
}

func Push(s []byte, x ...byte) []byte {
	return append(s, x...)
}

func PushBack(s []byte, x ...byte) []byte {
	return append([]byte(x), s...)
}
