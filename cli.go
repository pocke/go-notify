package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type CLI struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer

	Options *Options
}

type Options struct {
	Version bool
}

func NewCLI() *CLI {
	return &CLI{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
		Options:   &Options{},
	}
}

func (c *CLI) Run(args []string) int {
	err := c.ParseOption(args)
	if err != nil {
		return ExitCodeParseFlagError
	}

	if c.Options.Version {
		fmt.Fprintf(c.ErrStream, "%s version %s\n", args[0], Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}

func (c *CLI) ParseOption(args []string) error {
	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.SetOutput(c.ErrStream)
	f.BoolVar(&c.Options.Version, "version", false, "Print version information and quit")

	return f.Parse(args[1:])
}
