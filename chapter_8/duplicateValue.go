package chapter_8

// Exercise 2 - Chapter 8
// Complexity time: O(N)
func DuplicateValue(arr []string) string {
	mapString := map[string]bool{}

	for _, value := range arr {
		if mapString[value] {
			return value
		}
		mapString[value] = true
	}

	return ""
}
