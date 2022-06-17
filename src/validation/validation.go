package validation

import (
	"log"
	"regexp"
)

func RegexValidation(tag string) bool {
	result := false
	semver, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])(|\\-rc\\.[0-9])$", tag)

	if !semver {
		log.Fatalf("This %s is not a valid tag.", tag)
	} else {
		result = semver
	}
	return result

}

func EnvValidation(tag string) string {
	result := "dev"
	envProd, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])$", tag)
	envDev, _ := regexp.MatchString("^([0-9]+)(\\.[0-9]+)(\\.[0-9])(\\-rc\\.[0-9])$", tag)

	if envProd {
		result = "prod"
	} else if envDev {
		result = "dev"
	} else {
		result = "no match"
	}
	return result
}
