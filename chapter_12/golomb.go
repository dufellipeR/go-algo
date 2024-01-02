package chapter_12

func Golomb(n int, memo map[int]int) int {
	//fmt.Println("Call made")

	if n == 1 {
		return 1
	}

	if memo[n] == 0 {

		memo[n] = 1 + Golomb(n-Golomb(Golomb(n-1, memo), memo), memo)
	}

	return memo[n]
}

//func Golomb(n int, memo map[int]int) int {
//	fmt.Println("Call made")
//
//	if n == 1 {
//		return 1
//	}
//
//	memoization := memo[n]
//
//	if memoization != 0 {
//		fmt.Println("Enter here once")
//		return 1 + memoization
//	} else {
//		memo[n] = Golomb(n-1, memo)
//		return 1 + Golomb(n-Golomb(Golomb(n-1, memo), memo), memo)
//	}
//
//}

//func Golomb(n int) int {
//	if n == 1 {
//		return 1
//	}
//
//	return 1 + Golomb(n-Golomb(Golomb(n-1)))
//}
