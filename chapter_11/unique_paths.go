package chapter_11

import "fmt"

func UniquePaths(rows int, columns int) int {
	fmt.Println("made call")
	if rows == 1 || columns == 1 {
		return 1
	}

	return UniquePaths(rows-1, columns) + UniquePaths(rows, columns-1)
}
