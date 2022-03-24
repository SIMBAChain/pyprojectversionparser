package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPyProject(t *testing.T) {
	path := "../testfiles/pyproject.toml"

	parser := PyProjectDotToml{}

	details, err := parser.Parse(path)

	assert.Nil(t, err, "Parse thew an unexpected error")

	if details.Version != "PACKAGE_VERSION=1.2.3" {
		t.Errorf("PyProject parser found the wrong version [got: %s, expected: %s]", details.Version, "PACKAGE_VERSION=1.2.3")
	}

	if details.Name != "PACKAGE_NAME=a-py-package" {
		t.Errorf("PyProject parser found the wrong name [got: %s, expected: %s]", details.Name, "PACKAGE_NAME=a-py-package")
	}

}

func TestBrokenPyProject(t *testing.T) {
	path := "../testfiles/pyproject.broken.toml"

	parser := PyProjectDotToml{}

	details, err := parser.Parse(path)

	assert.Error(t, err, "Parse should throw an error")
	assert.Nil(t, details, "Details should be nil")
}
