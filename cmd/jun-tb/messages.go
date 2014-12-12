package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/sorcix/irc"
)

type MessageBuffer struct {
	buffer []*irc.Message
}

func NewMessageBuffer(capacity int) *MessageBuffer {
	return &MessageBuffer{make([]*irc.Message, 0, capacity)}
}

func (m *MessageBuffer) Len() int {
	return len(m.buffer)
}

func (m *MessageBuffer) Get(i int) *irc.Message {
	return m.buffer[i]
}

func (m *MessageBuffer) Add(message *irc.Message) {
	m.buffer = append(m.buffer, message)
}

type MessageView struct {
	buffer *MessageBuffer
	page   int
}

func NewMessageView(buffer *MessageBuffer) *MessageView {
	return &MessageView{buffer, 0}
}

func (m *MessageView) Draw(x, y, w, h int) {
	var (
		message    string
		lineHeight int
		ly         = h
	)

	for i := (m.buffer.Len() - 1); i >= 0; i-- {
		message = FormatMessage(m.buffer.Get(i))
		lineHeight = len(message)/w + 1
		ly -= lineHeight

		DrawString(message, x, ly, w, h)
	}
}

func (m *MessageView) PageUp() {
	// TODO
}

func (m *MessageView) PageDown() {
	// TODO
}

func FormatMessage(message *irc.Message) (out string) {
	switch message.Command {
	case irc.PRIVMSG:
		out = fmt.Sprintf("%s> %s", message.Name, message.Trailing)
	default:
		out = fmt.Sprintf("DEBUG: %s", message.String())
	}

	return
}

// TODO: Set color of cells.
func DrawString(s string, x, y, w, h int) {
	lx, ly := x, y

	for _, r := range s {
		if lx == w {
			lx, ly = x, ly+1
		}

		termbox.SetCell(lx, ly, r, termbox.ColorDefault, termbox.ColorDefault)
		lx++
	}
}
