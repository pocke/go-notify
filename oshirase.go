package main

import (
	"fmt"

	"github.com/pocke/oshirase"
)

var f = func(n *oshirase.NotifyArg) {
	fmt.Println(n.ID)
	fmt.Println(n.Summary)
	fmt.Println(n.Body)
}

func RunServer(name string) error {
	srv, err := oshirase.NewServer(name, "Pocke", Version)
	if err != nil {
		return err
	}

	srv.OnNotify(f)
	srv.OnCloseNotification(func(_ uint32) bool { return true })

	fmt.Println("start")

	return nil
}
