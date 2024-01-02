package tests_test

import (
	"algorithms/chapter_11"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EvenNumbers", func() {
	It("should test that the solution returns the correct value", func() {
		array1 := []int{0, 2, 3, 45, 42, 3, 21, 256}
		Expect(chapter_11.EvenNumbers(array1)).To(Equal([]int{0, 2, 42, 256}))
	})
})
