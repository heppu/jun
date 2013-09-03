package jun

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

type Jun struct {
	Address   string
	Nickname  string
	Channels  []string
	TlsConfig *tls.Config // http://golang.org/pkg/crypto/tls/#Config
	Quit      chan bool
	callbacks map[string][]Callback

	connection net.Conn
	send       chan string   // Messages to the server
	receive    chan *Message // Messages from the server
}

func New(address, nickname string, channels []string, tlsConfig *tls.Config) *Jun {
	return &Jun{
		Address:    address,
		Nickname:   nickname,
		Channels:   channels,
		TlsConfig:  tlsConfig,
		Quit:       make(chan bool),
		callbacks:  make(map[string][]Callback),
		connection: nil,
		send:       nil,
		receive:    nil,
	}
}

func (j *Jun) callbackLoop() {
	for message := range j.receive {
		if callbacks, ok := j.callbacks[message.Command]; ok {
			for _, cb := range callbacks {
				cb(message)
			}
		}
	}
}

func (j *Jun) receiveLoop() {
	var message *Message
	reader := bufio.NewReader(j.connection)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("\x1b[31m!!\x1b[0m %s\n", err)
			break
		}

		message = ParseMessage(line)
		log.Printf("\x1b[34m<<\x1b[0m %s\n", message)
		j.receive <- message
	}
}

func (j *Jun) sendLoop() {
	for data := range j.send {
		if _, err := j.connection.Write([]byte(data + "\r\n")); err != nil {
			log.Printf("\x1b[31m!!\x1b[0m %s\n", err)
		} else {
			log.Printf("\x1b[32m>>\x1b[0m %s\n", data)
		}
	}
}

func (j *Jun) Connect() (err error) {
	j.send = make(chan string, 32)
	j.receive = make(chan *Message, 32)

	if j.TlsConfig != nil {
		j.connection, err = tls.Dial("tcp", j.Address, j.TlsConfig)
	} else {
		j.connection, err = net.Dial("tcp", j.Address)
	}
	if err != nil {
		return
	}

	go j.sendLoop()
	go j.receiveLoop()
	go j.callbackLoop()

	j.AddCallback("266", j.raw266)
	j.AddCallback("433", j.nickInUse)
	j.AddCallback("PING", j.pingBack)
	j.AddCallback("NICK", j.nickChange)

	j.User(j.Nickname, j.Nickname)
	j.Nick(j.Nickname)

	return
}

func (j *Jun) Disconnect() {
	j.connection.Close()
	close(j.send)
	close(j.receive)
	j.Quit <- true
}

func (j *Jun) AddCallback(name string, cb Callback) {
	if _, ok := j.callbacks[name]; ok {
		j.callbacks[name] = append(j.callbacks[name], cb)
	} else {
		j.callbacks[name] = []Callback{cb}
	}
}

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
