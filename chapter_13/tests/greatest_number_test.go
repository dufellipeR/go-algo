package tests_test

import (
	"algorithms/chapter_13"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("GreatestNumber", func() {
	It("should correctly return the greatest number in the array", func() {
		arr := []int{3, 5, 7, 9, 4, 9, 4, 2, 21}
		gomega.Expect(chapter_13.ON(arr)).To(gomega.Equal(21))
		gomega.Expect(chapter_13.ONLogN(arr)).To(gomega.Equal(21))
		gomega.Expect(chapter_13.ON2(arr)).To(gomega.Equal(21))
	})
})
