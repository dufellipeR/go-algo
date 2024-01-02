package chapter_13

func MissingNumber(arr []int) int {
	sorted := QuickSortStart(arr)

	for i := 0; i < len(sorted); i++ {
		if !(sorted[i] == i) {
			return i
		}
	}
	return 0
}
