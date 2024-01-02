package kata

import "strings"

func Capitalize(st string) (capitalized [2]string) {
	var even string
	var odd string

	for key, char := range st {
		stringfy := string(char)
		if (key & 1) == 0 {
			even += strings.ToUpper(stringfy)
			odd += stringfy
			continue
		} else {
			even += strings.ToUpper(stringfy)
			odd += strings.ToUpper(stringfy)
		}
	}

	capitalized[0] = even
	capitalized[1] = odd

	return
}
