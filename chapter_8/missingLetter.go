package chapter_8

import (
	"strings"
)

// Exercise 3 - Chapter 8
// Complexity time: O(N)
func MissingLetter(word string) string {
	stringMap := map[rune]bool{}
	alphabetArr := []rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'x',
		'y',
		'z',
	}

	for _, value := range strings.Trim(word, " ") {
		stringMap[value] = true
	}

	for _, letter := range alphabetArr {
		if !stringMap[letter] {
			return string(letter)
		}
	}

	return ""
}
