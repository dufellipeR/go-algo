package chapter_11

func EvenNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	if numbers[0]%2 == 0 {
		return append([]int{numbers[0]}, EvenNumbers(numbers[1:])...)

	} else {
		return EvenNumbers(numbers[1:])
	}

}

//func EvenNumbers(numbers []int) []int {
//	var arr []int
//	if len(numbers) == 1 {
//
//		if numbers[0]%2 == 0 {
//			return []int{numbers[0]}
//		} else {
//			return []int{}
//		}
//
//	}
//
//	arr = append(arr, EvenNumbers(numbers[1:])...)
//
//	return arr
//}
