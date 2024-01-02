package chapter_10

import "fmt"

func Countdown(from int) {
	fmt.Println(from)
	if from == 0 {
		return
	} else {
		Countdown(from - 1)
	}
}
