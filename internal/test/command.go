package test

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type CommandRunner struct {
	command *exec.Cmd
}

func NewCommandRunner(ctx context.Context, cmd string, args ...string) *CommandRunner {
	toRun := exec.CommandContext(ctx, cmd, args...)
	return &CommandRunner{
		command: toRun,
	}
}

func (c *CommandRunner) WithStdIn(in io.Reader) *CommandRunner {
	c.command.Stdin = in
	return c
}

func (c *CommandRunner) Execute() (out string, err error) {
	// Commands referencing local executables will have path separators
	// Strip these out when creating temp files
	baseCmd := filepath.Base(c.command.Path)

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
	c.command.Stdout = combinedOut
	c.command.Stderr = combinedOut
	runErr := c.command.Run()
	outBytes, readErr := os.ReadFile(combinedOut.Name())
	out = string(outBytes)
	if runErr != nil {
		err = fmt.Errorf("command failed with exit code %d, output: %s",
			c.command.ProcessState.ExitCode(), out)
	}
	if readErr != nil && err == nil {
		err = readErr
		return
	}

	return
}
