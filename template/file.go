package template

import (
	"fmt"
	"os"
	tmpl "text/template"
)

var (
	funcMap map[string]interface{}
)

func init() {
	funcMap := make(map[string]interface{})
	funcMap["env"] = getEnv
}

func getEnv(envname string) (string, error) {
	env := os.Getenv(envname)
	if env == "" {
		return "", fmt.Errorf("Could not find the enviornment variable %s", envname)
	}
	return env, nil
}

func ProcessTemplate(file Template) error {
	// assume that file exists and is a tmpl file
	fmt.Printf("Processing env replacement for source file (%s) to target file (%s)\n", file.Name, file.TargetFile)
	confTmpl, err := tmpl.ParseFiles(file.Name)

	if err != nil {
		return err
	}

	confTmpl.Funcs(funcMap)

	err = os.MkdirAll(file.TargetFile, 0755)

	if err != nil {
		return err
	}

	outfile, err := os.Create(file.TargetFile)

	if err != nil {
		return err
	}

	err = confTmpl.Execute(outfile, nil)

	if err != nil {
		return err
	}

	err = outfile.Close()

	if err != nil {
		return err
	}

	return nil
}
