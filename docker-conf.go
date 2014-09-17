package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/docker-conf/template"
)

var (
	templateDir        = ""
	defaultTemplateDir = "/opt/docker-conf/templates"
	printVersion       bool
)

func init() {
	flag.StringVar(&templateDir, "template-dir", "", "directory containing docker-conf template files")
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
}

func main() {
	flag.Parse()

	if printVersion {
		fmt.Printf("docker-conf %s\n", Version)
		os.Exit(0)
	}

	if templateDir == "" {
		templateDir = defaultTemplateDir
	}
	if !IsFileExist(templateDir) {
		fmt.Printf("template directory (%s) does not exist\n", templateDir)
		os.Exit(1)
	}

	templates, err := template.WalkTemplateDir(templateDir)

	if err != nil {
		fmt.Printf("failed to walk template directory: %s\n", err.Error())
		os.Exit(1)
	}

	for _, element := range templates {
		fmt.Printf("Processing: %s\n", element.Name)
	}
}

// IsFileExist reports whether path exits.
func IsFileExist(fpath string) bool {
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false
	}
	return true
}
