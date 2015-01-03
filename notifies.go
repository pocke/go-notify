package main

import (
	"fmt"
	"sync"

	"github.com/pocke/oshirase"
)

type Notifies struct {
	notifyes []oshirase.Notify
	nmu      sync.RWMutex
}

func NewNotifies() *Notifies {
	return &Notifies{}
}

func (ns *Notifies) Add(n *oshirase.Notify) {
	ns.nmu.Lock()
	defer ns.nmu.Unlock()

	ns.notifyes = append(ns.notifyes, *n)
}

func (ns *Notifies) Delete(id uint32) error {
	idx, err := ns.FindByID(id)
	if err != nil {
		return err
	}

	ns.nmu.Lock()
	defer ns.nmu.Unlock()
	ns.notifyes = append(ns.notifyes[:idx], ns.notifyes[idx+1:]...)
	return nil
}

func (ns *Notifies) FindByID(id uint32) (index int, err error) {
	ns.nmu.RLock()
	defer ns.nmu.RUnlock()

	for i, n := range ns.notifyes {
		if n.ID == id {
			return i, nil
		}
	}

	return -1, fmt.Errorf("id %d dosesn't exist", id)
}
