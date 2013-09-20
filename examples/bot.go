package main

import (
	"crypto/tls"
	"github.com/FSX/jun/irc"
	"github.com/FSX/jun/plugins/pokemon"
	"os"
	"os/signal"
)

func main() {
	e := irc.New(
		"irc.rizon.net:9999",
		"Somerandomnickname",
		[]string{"#somerandomchannel"},
		&tls.Config{InsecureSkipVerify: true},
	)
	pokemon.PokemonQuotes(e)
	e.Connect()

	// Graceful shutdown for Ctrl+C
	go func(e *irc.Jun) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		e.Disconnect()
		os.Exit(0)
	}(e)

	<-e.Quit
}
