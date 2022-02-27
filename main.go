package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "remote", "--verbose", "|", "awk '{print $2}'", "|", "head", "-n", "1")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Print(string(stdout))
	}
}
