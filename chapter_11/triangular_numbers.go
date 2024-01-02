package chapter_11

func Triangular_numbers(n int) int {

	if n == 1 {
		return n
	}

	return Triangular_numbers(n-1) + n
}
