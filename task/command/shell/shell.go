package shell

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"time"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/2 22:25:30
* @description:
**/

type Command struct {
	bash           string
	instruction    string
	stdout, stderr *bytes.Buffer
}

func New(instruction string) *Command {
	return &Command{
		instruction: instruction,
		bash:        "/bin/bash",
		stdout:      &bytes.Buffer{},
		stderr:      &bytes.Buffer{},
	}
}

func (c *Command) SetBash(bash string) {
	c.bash = bash
}

func (c *Command) Run() error {
	cmd := exec.Command(c.bash, "-c", c.instruction)
	return cmd.Run()
}

func (c *Command) RunWait() error {
	cmd := exec.Command(c.bash, "-c", c.instruction)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *Command) RunWaitTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, c.bash, "-c", c.instruction)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return errors.New("command timed out")
		}
		return err
	}
	return nil
}

func (c *Command) GetStdout() *bytes.Buffer {
	return c.stdout
}

func (c *Command) GetStderr() *bytes.Buffer {
	return c.stderr
}
