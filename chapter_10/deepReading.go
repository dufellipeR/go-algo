package chapter_10

import "fmt"

// Exercise 4 - Chapter 10
// Complexity O(N)
// Learned about recursive functions and switch types with interfaces

func DeepReading(numbers []interface{}) {

	for _, value := range numbers {
		switch v := value.(type) {
		case int:
			fmt.Println(v)
		case []interface{}:
			DeepReading(v)
		}
	}
}
