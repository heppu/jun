package jun

import "fmt"

func (j *Jun) Raw(data string) {
	j.send <- data
}

func (j *Jun) User(username, realname string) {
	j.send <- fmt.Sprintf("USER %s +iw * :%s", username, realname)
}

func (j *Jun) Nick(nickname string) {
	j.send <- fmt.Sprintf("NICK %s", nickname)
}

func (j *Jun) Join(channel string) {
	j.send <- fmt.Sprintf("JOIN %s", channel)
}

func (j *Jun) Names(channel string) {
	j.send <- fmt.Sprintf("NAMES %s", channel)
}

func (j *Jun) Pong(payload string) {
	j.send <- fmt.Sprintf("PONG :%s", payload)
}

func (j *Jun) Privmsg(target, message string) {
	j.send <- fmt.Sprintf("PRIVMSG %s :%s", target, message)
}

func (j *Jun) QuitMsg() {
	j.send <- fmt.Sprintf("QUIT")
}
