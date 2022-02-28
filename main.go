package main

import (
	"fmt"
	"local/src/validation"
	"os"
)

func main() {

	tag := os.Getenv("GIT_REPO_TAG")

	matched := validation.RegexValidation(tag)
	environ := validation.EnvValidation(tag)

	fmt.Printf("This branch is ok: %t \nFor this environment: %s \n", matched, environ)
}
