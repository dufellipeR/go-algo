package chapter_8

// Exercise 1 - Chapter 8
// Complexity time: O(N)
func Intersection(arr1, arr2 []int) []int {
	var newArray []int
	var largeArray []int
	var smallArray []int
	mapArray := map[int]bool{}

	if len(arr1) > len(arr2) {
		largeArray = arr1
		smallArray = arr2
	} else {
		largeArray = arr2
		smallArray = arr1
	}

	for _, value := range largeArray {
		mapArray[value] = true
	}

	for _, value := range smallArray {
		if mapArray[value] {
			newArray = append(newArray, value)
		}
	}

	return newArray
}
