package main_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKidBase10(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kidbase10 CLI")
}

func ExecuteCommand(ctx context.Context, cmd string, args ...string) (out string, err error) {
	toRun := exec.Command(cmd, args...)
	// Commands referencing local executables will have path separators
	// Strip these out when creating temp files
	baseCmd := filepath.Base(cmd)

	combinedOut, err := os.CreateTemp("", fmt.Sprintf("%s-combined-out-*.log", baseCmd))
	if err != nil {
		return out, err
	}
	defer func() {
		rmErr := os.Remove(combinedOut.Name())
		if err == nil {
			err = rmErr
		}
	}()
	toRun.Stdout = combinedOut
	toRun.Stderr = combinedOut
	runErr := toRun.Run()
	outBytes, readErr := os.ReadFile(combinedOut.Name())
	out = string(outBytes)
	if runErr != nil {
		err = fmt.Errorf("command failed with exit code %d, output: %s",
			toRun.ProcessState.ExitCode(), out)
	}
	if readErr != nil && err == nil {
		err = readErr
		return
	}

	return
}

var _ = BeforeSuite(func(ctx context.Context) {

	By("building the encoder")
	Expect(ExecuteCommand(ctx, "make", "clean", "build")).Error().NotTo(HaveOccurred())

})
