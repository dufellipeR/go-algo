package tests_test

import (
	"algorithms/chapter_11"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountChars", func() {
	It("should test that the solution returns the correct value", func() {
		arrStrings := []string{"ab", "c", "def", "ghigergewrgj"}
		Expect(chapter_11.CountChars(arrStrings)).To(Equal(18))
	})
})
