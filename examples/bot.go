package main

import (
	"crypto/tls"
	"github.com/FSX/jun"
	"github.com/FSX/jun/plugins"
	"os"
	"os/signal"
)

func main() {
	e := jun.New(
		"irc.rizon.net:9999",
		"Somerandomnickname",
		[]string{"#somerandomchannel"},
		&tls.Config{InsecureSkipVerify: true},
	)
	plugins.PokemonQuotes(e)
	e.Connect()

	// Graceful shutdown for Ctrl+C
	go func(e *jun.Jun) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		e.Disconnect()
		os.Exit(0)
	}(e)

	<-e.Quit
}
