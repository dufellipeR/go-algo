package tests_test

import (
	"algorithms/chapter_13"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("MissingNumber", func() {
	It("should find the missing number in the array", func() {
		arr := []int{5, 2, 4, 1, 0}
		missingNumber := chapter_13.MissingNumber(arr)
		gomega.Expect(missingNumber).To(gomega.Equal(3))
	})
})
