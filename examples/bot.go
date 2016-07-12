package main

import (
	"crypto/tls"
	"os"
	"os/signal"

	"github.com/FSX/jun/client"
	"github.com/FSX/jun/plugins/pokemon"
)

func main() {
	e := client.New(
		"irc.rizon.net:9999",
		"Somerandomnickname",
		[]string{"#somerandomchannel"},
		&tls.Config{InsecureSkipVerify: true},
	)
	pokemon.PokemonQuotes(e)
	e.Connect()

	// Graceful shutdown for Ctrl+C
	go func(e *client.Client) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		e.Disconnect()
		os.Exit(0)
	}(e)

	<-e.Quit
}
