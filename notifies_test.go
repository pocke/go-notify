package main

import (
	"testing"

	"github.com/pocke/oshirase"
)

func TestNewNotifies(t *testing.T) {
	notifies := NewNotifies()
	if notifies == nil {
		t.Errorf("NewNotifies should return Notifies pointer. But got nil.")
	}
}

func TestNotifiesAdd(t *testing.T) {
	n := &oshirase.Notify{}
	ns := NewNotifies()

	l := len(ns.notifies)
	ns.Add(n)

	if l+1 != len(ns.notifies) {
		t.Errorf("Notifies.Add should add notify into notifies. But not Add.")
	}
}

func TestNotifesDelete(t *testing.T) {
	ns := NewNotifies()

	for i := 0; i < 10; i++ {
		n := &oshirase.Notify{
			ID: uint32(i),
		}
		ns.Add(n)
	}

	err := ns.Delete(uint32(5))
	if err != nil {
		t.Errorf("Expected: err is nil, Got: %s", err)
	}

	err = ns.Delete(uint32(5))
	if err == nil {
		t.Errorf("Expected: error, Got: %s", err)
	}

	if len(ns.notifies) != 9 {
		t.Errorf("notifies len, Expected: %d, Got: %d", 9, len(ns.notifies))
	}
}

func TestNotifiesFindByID(t *testing.T) {
	ns := NewNotifies()

	for i := 0; i < 10; i++ {
		n := &oshirase.Notify{
			ID: uint32(i),
		}
		ns.Add(n)
	}

	idx, err := ns.FindByID(uint32(1))
	if err != nil {
		t.Errorf("Expected: err is nil, Got: %s", err)
	}
	if idx < 0 {
		t.Errorf("idx should be 0 and more. But got %d", idx)
	}

	idx, err = ns.FindByID(uint32(1000))
	if err == nil {
		t.Errorf("When id doesn't exist, should return error. But got nil")
	}
	if idx >= 0 {
		t.Errorf("When id doesn't exist, index should be less than 0, but got %d", idx)
	}
}
