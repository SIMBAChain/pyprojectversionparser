package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageJson(t *testing.T) {
	path := "../testfiles/package.json"

	parser := PackageDotJson{}

	details, err := parser.Parse(path)

	assert.Nil(t, err, "Parse thew an unexpected error")

	if details.Version != "PACKAGE_VERSION=3.2.1" {
		t.Errorf("PyProject parser found the wrong version [got: %s, expected: %s]", details.Version, "PACKAGE_VERSION=3.2.1")
	}

	if details.Name != "PACKAGE_NAME=a-node-package" {
		t.Errorf("PyProject parser found the wrong name [got: %s, expected: %s]", details.Name, "PACKAGE_NAME=a-node-package")
	}

}

func TestBrokenPackageJson(t *testing.T) {
	path := "../testfiles/package.broken.json"

	parser := PackageDotJson{}

	details, err := parser.Parse(path)

	assert.Error(t, err, "Parse should throw an error")
	assert.Nil(t, details, "Details should be nil")
}
