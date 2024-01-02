package chapter_11

func FindX(str string) int {

	if str[0] == 'x' {
		return 0
	}

	return FindX(str[1:]) + 1
}
