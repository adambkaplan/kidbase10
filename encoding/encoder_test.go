package encoding

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encoder", Label("unit"), func() {

	When("default settings are used", func() {

		var (
			encoder *Encoder
			writer  *bytes.Buffer
		)

		BeforeEach(func() {
			writer = &bytes.Buffer{}
			encoder = NewEncoder(writer)
		})

		DescribeTable("samples should succeed",
			func(input string, output string) {
				Expect(encoder.Encode(input)).To(Succeed())
				Expect(writer.String()).To(Equal(output))
			},
			Entry("encodes 4-char string", "LANE", "51620000"),
			Entry("encodes < 8 characters", "I STARE", "40891720"),
			Entry("encodes 8 characters", "A STARER", "10891727"),
			Entry("encodes leading whitespace", "   LANE", "00051620"),
			Entry("encodes > 8 characters", "I STARE AT THE HILL", `40891720
19093203
45500000`),
		)
	})
})
