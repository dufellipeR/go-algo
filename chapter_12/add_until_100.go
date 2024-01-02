package chapter_12

// Complexity: O(N)

func AddUntil100(arr []int) int {
	//fmt.Println("Call made")
	if len(arr) == 0 {
		return 0
	}

	sum := AddUntil100(arr[1:])

	if arr[0]+sum > 100 {
		return sum
	} else {
		return arr[0] + sum
	}
}

// Complexity: O(2^N)

//func AddUntil100(arr []int) int {
//	fmt.Println("Call made")
//	if len(arr) == 0 {
//		return 0
//	}
//
//	if arr[0]+AddUntil100(arr[1:]) > 100 {
//		return AddUntil100(arr[1:])
//	} else {
//		return arr[0] + AddUntil100(arr[1:])
//	}
//}
