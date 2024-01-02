package chapter_13

func ON2(arr []int) int {
	var greatest int
	for _, ele := range arr {
		for _, value := range arr {
			if ele > value {
				greatest = ele
			}
		}
	}
	return greatest
}

func ONLogN(arr []int) int {
	sorted := QuickSortStart(arr)
	return sorted[len(sorted)-1]
}

func ON(arr []int) int {
	greatest := arr[0]

	for _, value := range arr {
		if value > greatest {
			greatest = value
		}
	}
	return greatest
}
