package client

import "fmt"

func (j *Client) Raw(data string) {
	j.send <- data
}

func (j *Client) User(username, realname string) {
	j.send <- fmt.Sprintf("USER %s +iw * :%s", username, realname)
}

func (j *Client) Nick(nickname string) {
	j.send <- fmt.Sprintf("NICK %s", nickname)
}

func (j *Client) Join(channel string) {
	j.send <- fmt.Sprintf("JOIN %s", channel)
}

func (j *Client) Names(channel string) {
	j.send <- fmt.Sprintf("NAMES %s", channel)
}

func (j *Client) Pong(payload string) {
	j.send <- fmt.Sprintf("PONG :%s", payload)
}

func (j *Client) Privmsg(target, message string) {
	j.send <- fmt.Sprintf("PRIVMSG %s :%s", target, message)
}

func (j *Client) QuitMsg() {
	j.send <- fmt.Sprintf("QUIT")
}
