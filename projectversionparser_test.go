package main

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
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
	defer func() {
		err = in.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		err = out.Close()
		if err != nil {
			log.Print(err)
		}
	}()

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
	err := Copy("testfiles/pyproject.toml", "./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"}

	main()

	err = os.Remove("./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
}

func TestMainNoArgsWithBadFile(t *testing.T) {
	clearOpts()
	err := Copy("testfiles/pyproject.broken.toml", "./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")

	err = os.Remove("./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
}

func TestMainNoArgsWithBadFileSetPath(t *testing.T) {
	clearOpts()
	err := Copy("testfiles/pyproject.broken.toml", "./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "./pyproject.toml"}

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", main, "os.Exit was not called")

	err = os.Remove("./pyproject.toml")
	if err != nil {
		log.Print(err)
		return
	}
}
