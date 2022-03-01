package main

import (
	"fmt"
	"local/src/command"
	"local/src/tekton"
	"local/src/validation"
	"os"
	"strconv"
	"strings"
)

func main() {

	tag := os.Args[1]

	gitElements := map[string]string{
		"commiter": "git log --pretty=format:\"%an: %ae\" --max-count=1",
		"commitId": "git rev-parse HEAD",
		"repoName": "git remote --verbose | awk '{print $2}' | head -n 1",
	}

	gitResults := map[string]string{
		"commiter":      "",
		"commitId":      "",
		"repoName":      "",
		"TagValidation": "",
		"Environment":   "",
	}

	gitResults["TagValidation"] = strconv.FormatBool(validation.RegexValidation(tag))
	gitResults["Environment"] = validation.EnvValidation(tag)

	for k, v := range gitElements {
		value := command.BashExecutor(v)
		tekton.ResultSender(k, strings.Replace(value, "\n", "", -1))
		gitResults[k] = value
	}

	for k, v := range gitResults {
		fmt.Printf("%s: %s \n", k, strings.Replace(v, "\n", "", -1))
	}
}
