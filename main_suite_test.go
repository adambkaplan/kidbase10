package main_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/kidbase10/internal/test"
)

func TestKidBase10(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kidbase10 CLI")
}

func ExecuteCommand(ctx context.Context, cmd string, args ...string) (string, error) {
	commandRunner := test.NewCommandRunner(ctx, cmd, args...)
	return commandRunner.Execute()
}

var _ = BeforeSuite(func(ctx context.Context) {

	By("building the encoder")
	Expect(ExecuteCommand(ctx, "make", "clean", "build")).Error().NotTo(HaveOccurred())

})
