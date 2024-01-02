package chapter_8

// Exercise 4 - Chapter 8
// Complexity time: O(N)
func NonDuplicateChar(word string) string {
	mapString := map[rune]int{}

	for _, value := range word {
		mapString[value] += 1
	}

	for key, _ := range mapString {
		if mapString[key] == 1 {
			return string(key)
		}
	}

	return ""
}
