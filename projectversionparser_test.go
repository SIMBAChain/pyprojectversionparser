package main

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMainPyProject(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/pyproject.toml"}
	main()
}

func TestMainPackageJson(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/package.json"}
	main()
}

func TestNoArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")
}

func TestMainDoesntExist(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/doesnt.exist"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")
}

func TestMainWithType(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-t", "pyproject.toml"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")
}

func TestMainWithTypeAndPath(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-t", "pyproject.toml", "testfiles/pyproject.toml"}

	main()
}
