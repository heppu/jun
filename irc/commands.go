package irc

import "fmt"

func (j *Bot) Raw(data string) {
	j.send <- data
}

func (j *Bot) User(username, realname string) {
	j.send <- fmt.Sprintf("USER %s +iw * :%s", username, realname)
}

func (j *Bot) Nick(nickname string) {
	j.send <- fmt.Sprintf("NICK %s", nickname)
}

func (j *Bot) Join(channel string) {
	j.send <- fmt.Sprintf("JOIN %s", channel)
}

func (j *Bot) Names(channel string) {
	j.send <- fmt.Sprintf("NAMES %s", channel)
}

func (j *Bot) Pong(payload string) {
	j.send <- fmt.Sprintf("PONG :%s", payload)
}

func (j *Bot) Privmsg(target, message string) {
	j.send <- fmt.Sprintf("PRIVMSG %s :%s", target, message)
}

func (j *Bot) QuitMsg() {
	j.send <- fmt.Sprintf("QUIT")
}
