package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	It("adds two numbers", func() {
		Expect(add(1, 1)).To(Equal(2))
	})
})
