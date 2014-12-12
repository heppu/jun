package main

import "github.com/nsf/termbox-go"

func termboxEvent(ev chan termbox.Event) {
	for {
		ev <- termbox.PollEvent()
	}
}

func initTermbox() (chan termbox.Event, error) {
	ev := make(chan termbox.Event)

	if err := termbox.Init(); err != nil {
		return nil, err
	}

	termbox.SetInputMode(termbox.InputEsc)
	go termboxEvent(ev)

	return ev, nil
}
