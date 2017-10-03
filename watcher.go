package main

import (
	"log"
	"time"

	watcher "github.com/radovskyb/watcher"
)

var watchInterval = time.Second * 5

func watch(filename string, updateFunc func()) {
	w := watcher.New()

	// Restrict to 1 event at a time.
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Move, watcher.Write, watcher.Create, watcher.Remove, watcher.Rename)

	go func() {
		for {
			select {
			case event := <-w.Event:
				log.Println("Changes to", event.Name())
				updateFunc()
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(filename); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(watchInterval); err != nil {
		log.Fatalln(err)
	}
}
