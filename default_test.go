package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("default encoding", Label("e2e", "conformance"), func() {

	When("valid text is passed with no arguments", func() {

		It("encodes text to an integer", func(ctx SpecContext) {
			Expect(ExecuteCommand(ctx, "./kidbase10", "LANE")).To(Equal("51620000"))
		})

		It("pads text less than 8 characters in length", func() {
			Skip("not implemented yet")
		})

		It("splits text longer than 8 characters into lines of 8", func() {
			Skip("not implented yet")
		})
	})
})
