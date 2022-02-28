package main

import (
	"fmt"
	"local/src/command"
	"local/src/validation"
	"os"
)

func main() {

	tag := os.Getenv("GIT_REPO_TAG")

	matched := validation.RegexValidation(tag)
	environ := validation.EnvValidation(tag)
	commiter := command.BashExecutor("git log --pretty=format:\"%an: %ae\" --max-count=1")
	commitId := command.BashExecutor("git rev-parse HEAD")
	repoName := command.BashExecutor("git remote --verbose | awk '{print $2}' | head -n 1")

	fmt.Printf("\nRepository: %s \nCommit ID: %s \nCommit Author: %s \n\nBranch is ok: %t \n\nEnvironment: %s \n",
		repoName, commitId, commiter, matched, environ)
}
