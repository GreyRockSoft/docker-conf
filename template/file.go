package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	tmpl "text/template"
)

func getFuncMap() map[string]interface{} {
	funcMap := make(map[string]interface{})
	funcMap["env"] = GetEnv

	return funcMap
}

func GetEnv(envname string) (string, error) {
	env := os.Getenv(envname)
	if env == "" {
		return "", fmt.Errorf("Could not find the enviornment variable %s", envname)
	}
	return env, nil
}

func ProcessTemplate(file Template) error {
	// assume that file exists and is a tmpl file
	fmt.Printf("Processing env replacement for source file (%s) to target file (%s)\n", file.Name, file.TargetFile)
	confTmpl := tmpl.New(file.TargetFile)

	fileContents, err := ioutil.ReadFile(file.Name)

	fmt.Printf("Adding custom functions and parsing template\n")
	confTmpl, err = confTmpl.Funcs(getFuncMap()).Parse(string(fileContents))

	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(file.TargetFile), 0755)

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
