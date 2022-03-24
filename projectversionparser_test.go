package main

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func clearOpts() {
	opts.Args.Path = ""
	opts.Type = ""
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func TestMainPyProject(t *testing.T) {
	clearOpts()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/pyproject.toml"}
	main()
}

func TestMainPackageJson(t *testing.T) {
	clearOpts()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/package.json"}
	main()
}

func TestNoArgs(t *testing.T) {
	clearOpts()
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
	clearOpts()
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
	clearOpts()
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

func TestMainWithBrokenFile(t *testing.T) {
	clearOpts()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "testfiles/pyproject.broken.toml"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")
}

func TestMainWithBrokenFileAndType(t *testing.T) {
	clearOpts()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-t", "pyproject.toml", "testfiles/pyproject.broken.toml"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")
}

func TestMainWithTypeAndPath(t *testing.T) {
	clearOpts()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-t", "pyproject.toml", "testfiles/pyproject.toml"}

	main()
}

func TestMainNoArgsWithFile(t *testing.T) {
	clearOpts()
	Copy("testfiles/pyproject.toml", "./pyproject.toml")
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"}

	main()

	os.Remove("./pyproject.toml")
}

func TestMainNoArgsWithBadFile(t *testing.T) {
	clearOpts()
	Copy("testfiles/pyproject.broken.toml", "./pyproject.toml")
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")

	os.Remove("./pyproject.toml")
}
