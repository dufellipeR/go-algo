package tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"algorithms/chapter_11"
)

var _ = Describe("TriangularNumbers", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(chapter_11.Triangular_numbers(7)).To(Equal(28))
	})
})
