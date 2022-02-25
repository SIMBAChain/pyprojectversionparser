package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
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

func main() {
	flag.Parse()

	values := flag.Args()

	if len(values) != 2 {
		fmt.Println("Usage: pyprojectversionparser tomlfile envfile")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(values[0]); err != nil {
		log.Fatal(err)
	}

	var poetry pyproject
	if _, err := toml.DecodeFile(values[0], &poetry); err != nil {
		log.Fatal(err)
	}

	ver := fmt.Sprintf("PACKAGE_VERSION=%s", poetry.Tool["poetry"].Version)
	name := fmt.Sprintf("PACKAGE_NAME=%s", poetry.Tool["poetry"].Name)

	fmt.Println(ver)
	fmt.Println(name)

	f, err := os.Create(values[1])
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.WriteString(ver); err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString("\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(name); err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString("\n"); err != nil {
		log.Fatal(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
