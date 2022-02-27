package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	tag := os.Getenv("GIT_REPO_TAG")

	matched := SemVer(tag)
	fmt.Print(matched)
}

func SemVer(a string) bool {

	result := false
	semverProd, _ := regexp.MatchString("^([0-9]+.[0-9.].[0-9])$", a)
	semverDev, _ := regexp.MatchString("^([0-9]+.[0-9.]+.[0-9]+-rc.[0-9])$", a)

	if semverProd {
		result = true
	} else if semverDev {
		result = true
	} else {
		result = false
	}

	return result
}
