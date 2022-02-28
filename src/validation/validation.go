package validation

import (
	"regexp"
)

func RegexValidation(a string) bool {
	result := false
	semver, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])(|\\-rc\\.[0-9])$", a)

	if semver {
		result = true
	} else {
		result = false
	}
	return result
}

func EnvValidation(a string) string {
	result := "dev"
	envProd, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])$", a)
	envDev, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])(\\-rc\\.[0-9])$", a)

	if envProd {
		result = "prod"
	} else if envDev {
		result = "dev"
	} else {
		result = "no match"
	}
	return result
}
