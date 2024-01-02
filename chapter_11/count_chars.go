package chapter_11

func CountChars(strings []string) int {

	if len(strings) == 1 {
		return len(strings[0])
	}

	return len(strings[0]) + CountChars(strings[1:])
}
