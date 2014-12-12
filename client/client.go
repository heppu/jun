package client

import (
	"bufio"
	"crypto/tls"
	// "log"
	"net"

	"github.com/sorcix/irc"
)

type Client struct {
	Address   string
	Nickname  string
	Channels  []string
	TlsConfig *tls.Config // http://golang.org/pkg/crypto/tls/#Config
	Quit      chan bool
	callbacks map[string][]Callback

	Error chan error

	connection net.Conn
	send       chan string       // Messages to the server
	receive    chan *irc.Message // Messages from the server
}

func New(address, nickname string, channels []string, tlsConfig *tls.Config) *Client {
	return &Client{
		Address:    address,
		Nickname:   nickname,
		Channels:   channels,
		TlsConfig:  tlsConfig,
		Quit:       make(chan bool),
		Error:      make(chan error, 32),
		callbacks:  make(map[string][]Callback),
		connection: nil,
		send:       nil,
		receive:    nil,
	}
}

func (j *Client) callbackLoop() {
	for message := range j.receive {
		if callbacks, ok := j.callbacks["*"]; ok {
			for _, cb := range callbacks {
				cb(message)
			}
		}
		if callbacks, ok := j.callbacks[message.Command]; ok {
			for _, cb := range callbacks {
				cb(message)
			}
		}
	}
}

func (j *Client) receiveLoop() {
	// var message *irc.Message
	reader := bufio.NewReader(j.connection)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			j.Error <- err
			break
		}

		j.receive <- irc.ParseMessage(line)
	}
}

func (j *Client) sendLoop() {
	for data := range j.send {
		if _, err := j.connection.Write([]byte(data + "\r\n")); err != nil {
			j.Error <- err
		}
	}
}

func (j *Client) Connect() (err error) {
	j.send = make(chan string, 32)
	j.receive = make(chan *irc.Message, 32)

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

func (j *Client) Disconnect() {
	j.connection.Close()
}
