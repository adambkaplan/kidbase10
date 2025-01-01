package encoding

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decoder", Label("unit", "decoding"), func() {

	When("default settings are used", func() {

		var (
			decoder *Decoder
			reader  *bytes.Buffer
		)

		BeforeEach(func() {
			reader = &bytes.Buffer{}
			decoder = NewDecoder(reader)
		})

		DescribeTable("samples should succeed",
			func(input string, output string) {
				Expect(reader.WriteString(input)).To(BeNumerically(">", 0))
				Expect(decoder.Decode()).To(BeEquivalentTo(output))
			},
			Entry("decodes with trailing zeros", "51620000", "LANE"),
			Entry("decodes < 8 characters", "40891720", "I STARE"),
			Entry("decodes 8 characters", "10891727", "A STARER"),
			Entry("decodes leading whitespace", "00051620", "   LANE"),
			Entry("decodes multi-line strings", `40891720
19093203
45500000`, "I STARE AT THE HILL"),
		)
	})
})
