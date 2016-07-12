package main

import "crypto/tls"

func main() {
	msgBuffer := NewMessageBuffer(2000)
	msgView := NewMessageView(msgBuffer)

	evQueue, err := initTermbox()
	if err != nil {
		panic(err)
	}

	cl, msgQueue, err := initIrc(
		"irc.zeronode.net:6697",
		"MrBleep",
		[]string{"#Hive6"},
		&tls.Config{InsecureSkipVerify: true},
	)

Loop:
	draw(msgView)

	select {
	case event := <-evQueue:
		switch event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc:
				goto Cleanup
			}
		case termbox.EventError:
			panic(event.Err)
		}
	case message := <-msgQueue:
		msgBuffer.Add(message)
	}

	goto Loop

Cleanup:
	cl.Disconnect()
	termbox.Close()
}

func draw(msgView *MessageView) {
	w, h := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	msgView.Draw(0, 0, w, h)
	termbox.Flush()
}
