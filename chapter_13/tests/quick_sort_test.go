package tests_test

import (
	"algorithms/chapter_13"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("QuickSort", func() {
	It("should sort an integer array", func() {
		arr := []int{3, 1, 2, 5, 6, 9, 7, 8, 4}
		chapter_13.QuickSortStart(arr)
		fmt.Println(arr)
		gomega.Expect(arr[0]).To(gomega.Equal(1))
		gomega.Expect(arr[len(arr)-1]).To(gomega.Equal(9))
	})
})
