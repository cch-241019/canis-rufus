package job

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"time"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/2 21:32:37
* @description:
**/

type Interface interface {
	Run() error
	RunWait() error
	RunWaitTimeout(time.Duration) error
	Stop() error
	GetStdout() *bytes.Buffer
	GetStderr() *bytes.Buffer
}

type Job struct {
	name    string
	cmdName string
	cmdArgs []string
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
}

func New(name string) *Job {
	return &Job{
		name:   name,
		stdout: new(bytes.Buffer),
		stderr: new(bytes.Buffer),
	}
}

func (job *Job) RunWait(instruction string) error {
	cmdName, cmdArgs, err := parseInstruction(instruction)
	if err != nil {
		return err
	}
	job.cmdName = cmdName
	job.cmdArgs = cmdArgs

	cmd := exec.Command(cmdName)

	cmd.Stdout = job.stdout
	cmd.Stderr = job.stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

var ErrInvalidCommand = errors.New("invalid command")

func parseInstruction(instruction string) (string, []string, error) {
	instruction = strings.TrimSpace(instruction)
	if len(instruction) == 0 {
		return "", []string{}, ErrInvalidCommand
	}
	cmdInfo := strings.Split(instruction, " ")
	if len(cmdInfo) == 1 {
		cmdName := strings.TrimSpace(cmdInfo[0])
		return cmdName, nil, nil
	}
	cmdName := strings.TrimSpace(cmdInfo[0])
	var cmdArgs []string
	for _, s := range cmdInfo[1:] {
		cmdArgs = append(cmdArgs, strings.TrimSpace(s))
	}
	return cmdName, cmdArgs, nil
}

func (job *Job) GetStdout() *bytes.Buffer {
	return job.stdout
}

func (job *Job) GetStderr() *bytes.Buffer {
	return job.stderr
}
