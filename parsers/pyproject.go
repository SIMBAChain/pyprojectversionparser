package parsers

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type (
	pyproject struct {
		Tool map[string]tool
	}

	tool struct {
		Name          string
		Version       string
		License       string
		Authors       []string
		Maintainers   []string
		Readme        string
		Homepage      string
		Repository    string
		Documentation string
		Keywords      []string
		Classifiers   []string
		Packages      []packages
		Include       []string
		Exclude       []string
	}
	packages struct {
		Include string
		From    string
		Format  string
	}
)

type PyProjectDotToml struct {
	IParser
}

func (p PyProjectDotToml) Parse(filename string) (*Details, error) {
	var poetry pyproject
	if _, err := toml.DecodeFile(filename, &poetry); err != nil {
		log.Print(err)
		return nil, err
	}

	details := Details{
		Version: fmt.Sprintf("PACKAGE_VERSION=%s", poetry.Tool["poetry"].Version),
		Name:    fmt.Sprintf("PACKAGE_NAME=%s", poetry.Tool["poetry"].Name),
	}

	return &details, nil
}
