package tests_test

import (
	"algorithms/chapter_11"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("../UniquePaths", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(chapter_11.UniquePaths(5, 5)).To(Equal(70))
	})
})
