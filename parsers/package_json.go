package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type packagejson struct {
	Name        string
	Version     string
	Description string
}

type PackageDotJson struct {
	IParser
}

func (p PackageDotJson) Parse(filename string) (*Details, error) {
	var data packagejson

	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Print(err)
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	details := Details{
		Version: fmt.Sprintf("PACKAGE_VERSION=%s", data.Version),
		Name:    fmt.Sprintf("PACKAGE_NAME=%s", data.Name),
	}

	return &details, nil
}
