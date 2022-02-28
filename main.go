package main

import (
	"fmt"
	"local/src/inventory"
	"local/src/validation"
	"os"
)

func main() {

	tag := os.Getenv("GIT_REPO_TAG")

	matched := validation.RegexValidation(tag)
	environ := validation.EnvValidation(tag)
	commiter := inventory.CommitAuthor()
	commitId := inventory.CommitId()
	repoName := inventory.RepoName()

	fmt.Printf("\nRepository: %s \nCommit ID: %s \nCommit Author: %s \n\nBranch is ok: %t \n\nEnvironment: %s \n",
		repoName, commitId, commiter, matched, environ)
}
