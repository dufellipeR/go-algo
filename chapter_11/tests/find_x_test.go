package tests_test

import (
	"algorithms/chapter_11"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindX", func() {
	It("should test that the solution returns the correct value", func() {
		str := "A purely peer-to-peerxversion of electronic cash would allow online payments to be sent directly from one party to another without going through a financial institution"
		Expect(chapter_11.FindX(str)).To(Equal(21))
	})
})
