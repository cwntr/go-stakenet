package common

import (
	"os/exec"
)

func ExecCLI(command string, arguments ...string) (string, error) {
	cmd := exec.Command(command, arguments...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

