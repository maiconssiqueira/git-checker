package main

import (
	"fmt"
	"https://github.com/maiconssiqueira/git-checker/main/src/command"
	"https://github.com/maiconssiqueira/git-checker/main/src/tekton"
	"https://github.com/maiconssiqueira/git-checker/main/src/command/validation"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	tag := ""

	if len(os.Args) > 1 {
		tag = os.Args[1]
	} else {
		log.Fatal("Need a valid tag to continue. Try including a tag in the excution arguments.")
	}

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
