package tests_test

import (
	"algorithms/chapter_13"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("GreatestProduct", func() {
	It("should return the gratest products of any 3 elements", func() {
		arr := []int{3, 1, 2, 5, 6, 9, 7, 8, 4}
		sorted := chapter_13.QuickSortStart(arr)
		greatest := chapter_13.GreatestProduct(sorted)
		gomega.Expect(greatest).To(gomega.Equal(504))
	})
})
