package relay

import (
	"fmt"
	"github.com/FSX/jun/irc"
	"strings"
)

func RelayChannels(a *irc.Bot, b *irc.Bot, serverA, channelA, serverB, channelB string) {
	AtoB := make(chan string)
	BtoA := make(chan string)
	relayOn := true

	makeCallback := func(relay chan string, server, channel string) irc.Callback {
		return func(message *irc.Message) {
			if channel != message.Arguments[0] {
				return
			}

			if strings.HasPrefix(message.Final, "!relay on") {
				relayOn = true
				relay <- "Relay is on."
			} else if strings.HasPrefix(message.Final, "!relay off") {
				relayOn = false
				relay <- "Relay is off."
			} else if relayOn {
				nick, _, _ := irc.ParsePrefix(message.Prefix)
				relay <- fmt.Sprintf("<%s@%s/%s> %s", nick, server, channel, message.Final)
			}
		}
	}

	a.AddCallback("PRIVMSG", makeCallback(AtoB, serverA, channelA))
	b.AddCallback("PRIVMSG", makeCallback(BtoA, serverB, channelB))

	go func(a, b *irc.Bot, channelA, channelB string) {
		for {
			select {
			case m := <-AtoB:
				b.Privmsg(channelB, m)
			case m := <-BtoA:
				a.Privmsg(channelA, m)
			case <-a.Quit:
				break
			case <-b.Quit:
				break
			}
		}
	}(a, b, channelA, channelB)
}
