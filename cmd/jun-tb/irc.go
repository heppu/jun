package main

import (
	"crypto/tls"

	"github.com/FSX/jun/client"
	"github.com/sorcix/irc"
)

func initIrc(
	address,
	nickname string,
	channels []string,
	tlsConfig *tls.Config,
) (*client.Client, chan *irc.Message, error) {
	cl := client.New(address, nickname, channels, tlsConfig)
	ch := make(chan *irc.Message)

	cl.AddCallback("*", func(message *irc.Message) {
		ch <- message
	})

	if err := cl.Connect(); err != nil {
		return nil, nil, err
	}

	return cl, ch, nil
}
