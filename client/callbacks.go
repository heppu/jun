package client

import "github.com/sorcix/irc"

type Callback func(*irc.Message)

func (j *Client) AddCallback(name string, cb Callback) {
	if _, ok := j.callbacks[name]; ok {
		j.callbacks[name] = append(j.callbacks[name], cb)
	} else {
		j.callbacks[name] = []Callback{cb}
	}
}

func (j *Client) raw266(message *irc.Message) {
	for _, channel := range j.Channels {
		j.Join(channel)
	}
}

func (j *Client) pingBack(message *irc.Message) {
	j.Pong(message.Trailing)
}

func (j *Client) nickInUse(message *irc.Message) {
	j.Nickname = j.Nickname + "_"
	j.Nick(j.Nickname)
}

func (j *Client) nickChange(message *irc.Message) {
	j.Nickname = message.Trailing
}
