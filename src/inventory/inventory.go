package inventory

import (
	"os/exec"
)

func RepoName() string {
	repoName := "git remote --verbose | awk '{print $2}' | head -n 1"
	cmd := exec.Command("/bin/sh", "-c", repoName)
	stdout, _ := cmd.Output()
	return string(stdout)
}

func CommitId() string {
	commitId := "git rev-parse HEAD"
	cmd := exec.Command("/bin/sh", "-c", commitId)
	stdout, _ := cmd.Output()
	return string(stdout)
}

func CommitAuthor() string {
	commitAuthor := "git log --pretty=format:\"%an: %ae\" --max-count=1"
	cmd := exec.Command("/bin/sh", "-c", commitAuthor)
	stdout, _ := cmd.Output()
	return string(stdout)
}
