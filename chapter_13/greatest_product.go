package chapter_13

func GreatestProduct(arr []int) int {
	sorted := QuickSortStart(arr)

	return sorted[len(sorted)-1] * sorted[len(sorted)-2] * sorted[len(sorted)-3]
}
