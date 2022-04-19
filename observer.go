package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Publisher interface {
	register(subscriber *Subscriber)
	unregister(subscriber *Subscriber)
	notify(path, event string)
	observe()
}

type Subscriber interface {
	receive(path, event string)
}

type PathWatcher struct {
	subscribers []*Subscriber
	watcher     fsnotify.Watcher
	rootPath    string
}

func (pw *PathWatcher) register(subscriber *Subscriber) {
	pw.subscribers = append(pw.subscribers, subscriber)
}

func (pw *PathWatcher) unregister(subscriber *Subscriber) {
	length := len(pw.subscribers)

	for i, sub := range pw.subscribers {
		if sub == subscriber {
			pw.subscribers[i] = pw.subscribers[length-1]
			pw.subscribers = pw.subscribers[:length-1]
			break
		}
	}
}

func (pw *PathWatcher) notify(path, event string) {
	for _, sub := range pw.subscribers {
		(*sub).receive(path, event)
	}
}

func (pw *PathWatcher) observe() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error", err)
	}
	defer watcher.Close()

	if err := filepath.Walk(pw.rootPath, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return watcher.Add(path)
		}

		return nil
	}); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				pw.notify(event.Name, event.Op.String())
			case err := <-watcher.Errors:
				fmt.Println("Error", err)
			}
		}
	}()

	<-done
}
