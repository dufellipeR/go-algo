package tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChapter11(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chapter11 Suite")
}
