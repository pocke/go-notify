package main

import "os"

func main() {
	cli := NewCLI()
	code := cli.Run(os.Args)

	if !cli.Server {
		os.Exit(code)
	}

	select {}
}
