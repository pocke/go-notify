package main

import "os"

func main() {
	cli := NewCLI()
	cli.Run(os.Args)
}
