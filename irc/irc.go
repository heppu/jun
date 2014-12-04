package irc

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"

	xirc "github.com/sorcix/irc"
)

type Bot struct {
	Address   string
	Nickname  string
	Channels  []string
	TlsConfig *tls.Config // http://golang.org/pkg/crypto/tls/#Config
	Quit      chan bool
	callbacks map[string][]Callback

	connection net.Conn
	send       chan string        // Messages to the server
	receive    chan *xirc.Message // Messages from the server
}

func New(address, nickname string, channels []string, tlsConfig *tls.Config) *Bot {
	return &Bot{
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

func (j *Bot) callbackLoop() {
	for message := range j.receive {
		if callbacks, ok := j.callbacks[message.Command]; ok {
			for _, cb := range callbacks {
				cb(message)
			}
		}
	}
}

func (j *Bot) receiveLoop() {
	var message *xirc.Message
	reader := bufio.NewReader(j.connection)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("\x1b[31m!!\x1b[0m %s\n", err)
			break
		}

		message = xirc.ParseMessage(line)
		if err != nil {
			log.Printf("\x1b[31m!!\x1b[0m %s\n", err)
			break
		}

		log.Printf("\x1b[34m<<\x1b[0m %s\n", message)
		j.receive <- message
	}
}

func (j *Bot) sendLoop() {
	for data := range j.send {
		if _, err := j.connection.Write([]byte(data + "\r\n")); err != nil {
			log.Printf("\x1b[31m!!\x1b[0m %s\n", err)
		} else {
			log.Printf("\x1b[32m>>\x1b[0m %s\n", data)
		}
	}
}

func (j *Bot) Connect() (err error) {
	j.send = make(chan string, 32)
	j.receive = make(chan *xirc.Message, 32)

	j.AddCallback("266", j.raw266)
	j.AddCallback("433", j.nickInUse)
	j.AddCallback("PING", j.pingBack)
	j.AddCallback("NICK", j.nickChange)

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

	j.User(j.Nickname, j.Nickname)
	j.Nick(j.Nickname)

	return
}

func (j *Bot) Disconnect() {
	j.connection.Close()
	close(j.send)
	close(j.receive)
	j.Quit <- true
}
