package irc

import xirc "github.com/sorcix/irc"

type Callback func(*xirc.Message)

func (j *Bot) AddCallback(name string, cb Callback) {
	if _, ok := j.callbacks[name]; ok {
		j.callbacks[name] = append(j.callbacks[name], cb)
	} else {
		j.callbacks[name] = []Callback{cb}
	}
}

func (j *Bot) raw266(message *xirc.Message) {
	for _, channel := range j.Channels {
		j.Join(channel)
	}
}

func (j *Bot) pingBack(message *xirc.Message) {
	j.Pong(message.Trailing)
}

func (j *Bot) nickInUse(message *xirc.Message) {
	j.Nickname = j.Nickname + "_"
	j.Nick(j.Nickname)
}

func (j *Bot) nickChange(message *xirc.Message) {
	j.Nickname = message.Trailing
}
