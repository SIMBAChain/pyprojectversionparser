package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
	"pyprojectversionparser/parsers"
)

var opts struct {
	Type string `short:"t" long:"type" description:"Project type to parse" choice:"pyproject.toml" choice:"package.json"`
	// Example of positional arguments
	Args struct {
		Path string `positional-arg-name:"path/to/project/file"`
	} `positional-args:"yes"`
}

var parserMap = map[string]parsers.IParser{
	"pyproject.toml": parsers.PyProjectDotToml{},
	"package.json":   parsers.PackageDotJson{},
}

func main() {
	var err error

	_, err = flags.Parse(&opts)
	if err != nil {
		log.Print(err)
		os.Exit(-2)
	}

	var details *parsers.Details

	if opts.Args.Path == "" && opts.Type != "" {
		details, err = parserMap[opts.Type].Parse(opts.Type)
		if err != nil {
			os.Exit(-1)
		}
	} else if opts.Args.Path != "" && opts.Type != "" {
		details, err = parserMap[opts.Type].Parse(opts.Args.Path)
		if err != nil {
			os.Exit(-1)
		}
	} else if opts.Args.Path != "" && opts.Type == "" {
		file := filepath.Base(opts.Args.Path)
		if val, ok := parserMap[file]; ok {
			details, err = val.Parse(opts.Args.Path)
			if err != nil {
				os.Exit(-1)
			}
		} else {
			log.Printf("Can't find parser for %s", file)
			os.Exit(-1)
		}
	} else {
		var done = false
		for key, val := range parserMap {
			if _, err := os.Stat(key); err == nil {
				details, err = val.Parse(key)
				if err != nil {
					os.Exit(-1)
				}
				done = true
				break
			}
		}
		if !done {
			log.Print("No valid project files found!")
			os.Exit(-1)
		}
	}

	fmt.Println(details.Version)
	fmt.Println(details.Name)
}
