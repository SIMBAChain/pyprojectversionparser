package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pyprojectversionparser/parsers"
)

var parserMap = map[string]parsers.IParser{
	"pyproject.toml": parsers.PyProjectDotToml{},
	"package.json":   parsers.PackageDotJson{},
}

func main() {
	flag.Parse()

	values := flag.Args()

	var details *parsers.Details

	if len(values) == 0 {
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
	} else if len(values) == 1 {
		file := filepath.Base(values[0])
		if _, err := os.Stat(values[0]); err != nil {
			log.Print(err)
			os.Exit(-1)
		}
		parser := parserMap[file]
		var err error
		details, err = parser.Parse(values[0])
		if err != nil {
			os.Exit(-1)
		}
	} else {
		fmt.Println("Usage: projectversionparser parsefile")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(details.Version)
	fmt.Println(details.Name)
}
