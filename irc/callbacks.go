package irc

type Callback func(*Message)

func (j *Bot) AddCallback(name string, cb Callback) {
	if _, ok := j.callbacks[name]; ok {
		j.callbacks[name] = append(j.callbacks[name], cb)
	} else {
		j.callbacks[name] = []Callback{cb}
	}
}

func (j *Bot) raw266(message *Message) {
	for _, channel := range j.Channels {
		j.Join(channel)
	}
}

func (j *Bot) pingBack(message *Message) {
	j.Pong(message.Final)
}

func (j *Bot) nickInUse(message *Message) {
	j.Nickname = j.Nickname + "_"
	j.Nick(j.Nickname)
}

func (j *Bot) nickChange(message *Message) {
	j.Nickname = message.Final
}
