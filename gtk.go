package main

import (
	"fmt"
	"time"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/pocke/oshirase"
)

func gthread(f func()) {
	gdk.ThreadsEnter()
	defer gdk.ThreadsLeave()
	f()
}

func StartGtk(appName string) {
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gtk.Init(&[]string{appName})

	gthread(func() {
		icon := gtk.NewStatusIcon()
		icon.SetName("go-notify")
		icon.SetFromFile("/home/pocke/Downloads/about_me/images/arch.png")
		icon.SetTooltipText(appName)
	})

	events.On("Notify", func(n *oshirase.Notify) {
		var popWin *gtk.Window
		gthread(func() {
			popWin = gtk.NewWindow(gtk.WINDOW_POPUP)
			popWin.SetSizeRequest(300, 50)
		})

		gthread(func() {
			flame := gtk.NewFrame(n.Summary)
			popWin.Add(flame)

			label := gtk.NewLabel(fmt.Sprintf("ID: %d, Body: %s", n.ID, n.Body))
			flame.Add(label)
		})

		gthread(func() {
			popWin.ShowAll()
		})

		time.Sleep(2 * time.Second)
		gthread(func() {
			popWin.Destroy()
		})
	})

	gthread(func() {
		gtk.Main()
	})
}
