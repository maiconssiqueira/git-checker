package command

import (
	"os/exec"
)

func BashExecutor(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, _ := cmd.Output()
	return string(stdout)
}
