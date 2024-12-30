package main_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func runEncoder(ctx context.Context, toEncode string) (string, error) {
	return ExecuteCommand(ctx, "./kidbase10", toEncode)
}

func runDecoder(ctx context.Context, toDecode string) (string, error) {
	return ExecuteCommand(ctx, "./kidbase10", "--decode", toDecode)
}

var _ = Describe("default encoding", Label("e2e", "conformance", "encoding"), func() {

	When("valid text is passed with no arguments", func() {

		It("encodes text to an integer", func(ctx SpecContext) {
			Expect(runEncoder(ctx, "LANE")).To(Equal("51620000"))
		})

		It("pads text less than 8 characters in length", func(ctx SpecContext) {
			Expect(runEncoder(ctx, "I STARE")).To(Equal("40891720"))
		})

		It("splits text longer than 8 characters into lines of 8", func(ctx SpecContext) {
			expectedOutput := `40891720
19093203
45500000`
			Expect(runEncoder(ctx, "I STARE AT THE HILL")).To(Equal(expectedOutput))
		})

		It("encodes leading whitespace as zeros", func(ctx SpecContext) {
			Expect(runEncoder(ctx, "   LANE")).To(Equal("00051620"))
		})
	})
})

var _ = Describe("default decoding", Label("e2e", "conformance", "decoding"), func() {

	When("valid encoded text is passed with the decode flag", func() {

		It("decodes an integer to text", func(ctx SpecContext) {
			Expect(runDecoder(ctx, "51620000")).To(Equal("LANE"))
		})

		It("trims trailing whitespace", func(ctx SpecContext) {
			Expect(runDecoder(ctx, "40891720")).To(Equal("I STARE"))
		})

		It("decodes multiline text", func(ctx SpecContext) {
			encodedInput := `40891720
19093203
45500000`
			Expect(runDecoder(ctx, encodedInput)).To(Equal("I STARE AT THE HILL"))
		})

		It("decodes leading zeros as whitespace", func(ctx SpecContext) {
			Expect(runDecoder(ctx, "00051620")).To(Equal("   LANE"))
		})
	})
})
