package encoding

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEncoding(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "encoding")
}
