package chapter_12

import "fmt"

func UniquePaths(rows int, columns int, memo map[string]int) int {
	fmt.Println("fmade call")
	key := fmt.Sprintf("%d_%d", rows, columns)
	if rows == 1 || columns == 1 {
		return 1
	}

	if memo[key] == 0 {
		memo[key] = UniquePaths(rows-1, columns, memo) + UniquePaths(rows, columns-1, memo)
	}

	return memo[key]
}
