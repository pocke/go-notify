package main

import (
	"fmt"

	"github.com/pocke/oshirase"
)

func RunServer(name string) error {
	srv, err := oshirase.NewServer(name, "Pocke", Version)
	if err != nil {
		return err
	}

	notifyes := NewNotifies()

	srv.OnNotify(func(n *oshirase.Notify) {
		notifyes.Add(n)
		fmt.Println(n.ID)
		fmt.Println(n.Summary)
		fmt.Println(n.Body)
	})
	srv.OnCloseNotification(func(id uint32) bool {
		err := notifyes.Delete(id)
		if err != nil {
			return false
		}
		return true
	})

	fmt.Println("start")

	return nil
}
