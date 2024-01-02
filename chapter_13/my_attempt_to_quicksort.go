package chapter_13

func QuickSortWrong(leftIndex int, rightIndex int, arr *[]int) {

	if rightIndex-leftIndex <= 0 {
		return
	}

	pivotIndex := partitionWrong(leftIndex, rightIndex, arr)

	QuickSortWrong(leftIndex, pivotIndex-1, arr)

	QuickSortWrong(pivotIndex+1, rightIndex, arr)
}

func partitionWrong(leftPointer int, rightPointer int, arr *[]int) int {

	pivotIndex := rightPointer

	pivot := (*arr)[pivotIndex]

	rightPointer -= 1

	for {
		for (*arr)[leftPointer] < pivot {
			leftPointer += 1
		}
		for (*arr)[rightPointer] > pivot {
			rightPointer -= 1
		}

		if leftPointer >= rightPointer {
			break
		} else {
			(*arr)[leftPointer], (*arr)[rightPointer] = (*arr)[rightPointer], (*arr)[leftPointer]
			leftPointer += 1
		}
	}

	return pivotIndex

}
