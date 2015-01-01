package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewCLI(t *testing.T) {
	cli := NewCLI()
	if cli.InStream != os.Stdin {
		t.Errorf("cli.InStream should be Stdin. But got %s", cli.InStream)
	}

	if cli.OutStream != os.Stdout {
		t.Errorf("cli.OutStream should be Stdout. But got %s", cli.OutStream)
	}

	if cli.ErrStream != os.Stderr {
		t.Errorf("cli.ErrStream should be Stdin. But got %s", cli.ErrStream)
	}

	if cli.Options == nil {
		t.Errorf("cli.Options should not be nil. But got nil")
	}
}

func TestRun(t *testing.T) {
	c, _, _, _ := newCLIForTest()
	code := c.Run([]string{"go-notify"})
	if code != ExitCodeOK {
		t.Errorf("Expected %d, But got %d", ExitCodeOK, code)
	}
}

func TestRunShowVersion(t *testing.T) {
	c, _, _, er := newCLIForTest()
	code := c.Run([]string{"go-notify", "--version"})
	if code != ExitCodeOK {
		t.Errorf("Expected %d, But got %d", ExitCodeOK, code)
	}
	if !strings.Contains(er.String(), Version) {
		t.Errorf("Output should include %s, but not.", Version)
		t.Errorf("Output: %q", er.String())
	}
}

func TestParseOption(t *testing.T) {
	c, _, _, _ := newCLIForTest()
	c.ParseOption([]string{"go-notify", "--version"})

	if !c.Options.Version {
		t.Errorf("When --version option, Options.Version should be true. But got false")
	}

	c, _, _, _ = newCLIForTest()
	c.ParseOption([]string{"go-notify", "-v"})

	if !c.Options.Version {
		t.Errorf("When -v option, Options.Version should be true. But got false")
	}
}

func newCLIForTest() (c *CLI, in, out, err *bytes.Buffer) {
	in = new(bytes.Buffer)
	out = new(bytes.Buffer)
	err = new(bytes.Buffer)

	c = &CLI{
		InStream:  in,
		OutStream: out,
		ErrStream: err,
		Options:   &Options{},
	}

	return
}
