package template

import (
	"fmt"
)

func ProcessTemplate(file Template) error {
	//assume that file exists and is a tmpl file
	fmt.Printf("Processing env replacement for source file (%s) to target file (%s)\n", file.Name, file.TargetFile)
	return nil
}
