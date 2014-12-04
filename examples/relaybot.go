package main

import (
	"crypto/tls"
	"github.com/FSX/jun/client"
	"github.com/FSX/jun/plugins/relay"
	"os"
	"os/signal"
)

func main() {

	Rizon := client.New(
		"irc.rizon.net:9999",
		"Somerandomnickname",
		[]string{"#somerandomchannel"},
		&tls.Config{InsecureSkipVerify: true},
	)
	Zeronode := client.New(
		"chat.freenode.net:6697",
		"Somerandomnickname",
		[]string{"#somerandomchannel"},
		&tls.Config{InsecureSkipVerify: true},
	)

	// Set up mirror
	relay.RelayChannels(Rizon, Zeronode,
		"R", "#somerandomchannel",
		"Z", "#somerandomchannel")

	// Connect
	Rizon.Connect()
	Zeronode.Connect()

	// Graceful shutdown for Ctrl+C
	go func(a *client.Client) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		Rizon.Disconnect()
		Zeronode.Disconnect()
		os.Exit(0)
	}(Rizon)

	<-Rizon.Quit
	<-Zeronode.Quit
}
