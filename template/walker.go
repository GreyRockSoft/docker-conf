package template

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	fileType = ".tmpl"
)

type Template struct {
	Name       string
	TargetFile string
}

func WalkTemplateDir(dir string) ([]Template, error) {
	templates := make([]Template, 0, 5)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, fileType) {
			templates = append(templates, Template{path, RemoveBasePath(dir, path)})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func RemoveBasePath(basePath, fullPath string) string {
	return fullPath
}
