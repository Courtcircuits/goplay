package mod

import (
	"os/exec"
)

// Exec executes a command and returns the output as a string
func Exec(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdoud, err := cmd.Output()
	return string(stdoud), err
}

func ExecSilent(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}
