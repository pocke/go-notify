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
	ExitCodeServerFailed
)

type CLI struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer

	Options *Options
	Server  bool
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
	name := args[0]

	err := c.ParseOption(args)
	if err != nil {
		return ExitCodeParseFlagError
	}

	if c.Options.Version {
		fmt.Fprintf(c.ErrStream, "%s version %s\n", name, Version)
		return ExitCodeOK
	}

	err = RunServer(name)
	if err != nil {
		return ExitCodeServerFailed
	}
	c.Server = true

	return ExitCodeOK
}

func (c *CLI) ParseOption(args []string) error {
	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.SetOutput(c.ErrStream)
	f.BoolVar(&c.Options.Version, "version", false, "Print version information and quit")

	return f.Parse(args[1:])
}
