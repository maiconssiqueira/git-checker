package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	cmd := "git tag --sort=taggerdate | tail -1"
	tag := exec.Command("bash", "-c", cmd)
	stdout, err := tag.Output()
	tagParse := string(stdout)
	tagParse = strings.TrimSuffix(tagParse, "\n")

	if err != nil {
		log.Fatal(err)
	}

	semverProd, _ := regexp.MatchString("^([0-9]+.[0-9.].[0-9])$", tagParse)
	semverDev, _ := regexp.MatchString("^([0-9]+.[0-9.]+.[0-9]-rc.[0-9])$", tagParse)

	if semverProd == true {
		fmt.Printf("Deu bom: é prod: %s \n", tagParse)
	} else if semverDev == true {
		fmt.Printf("Deu bom: é dev: %s \n", tagParse)
	} else {
		fmt.Printf("nao é nenhum, oq veio foi esse: %s \n", tagParse)
	}
}
