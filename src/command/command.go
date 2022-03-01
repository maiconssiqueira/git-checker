package command

import (
	"fmt"
	"os/exec"
)

func BashExecutor(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(stdout)
}
