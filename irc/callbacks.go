package irc

type Callback func(*Message)

func (j *Jun) raw266(message *Message) {
	for _, channel := range j.Channels {
		j.Join(channel)
	}
}

func (j *Jun) pingBack(message *Message) {
	j.Pong(message.Final)
}

func (j *Jun) nickInUse(message *Message) {
	j.Nickname = j.Nickname + "_"
	j.Nick(j.Nickname)
}

func (j *Jun) nickChange(message *Message) {
	j.Nickname = message.Final
}
