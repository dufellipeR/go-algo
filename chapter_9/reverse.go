package chapter_9

import (
	"fmt"
	"strings"
)

// Exercise 4 - Chapter 9
// Complexity time: O(N)

func Reverse(str string) string {
	stack := Stack{}

	var reversed string

	var charArr []string

	for _, char := range str {
		stack.Push(char)
	}

	for range str {
		char := stack.Pop()
		charArr = append(charArr, fmt.Sprintf("%c", char))
	}

	reversed = strings.Join(charArr, "")

	return reversed

}
