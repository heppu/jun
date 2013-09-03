package jun

import (
	"fmt"
	"strings"
)

type Message struct {
	Prefix    string
	Command   string
	Arguments []string
	Final     string
}

func (m *Message) String() string {
	// Could be better...
	return fmt.Sprintf("%s %s %s :%s", m.Prefix, m.Command, m.Arguments, m.Final)
}

// [:prefix ]command [arguments ]*[ :final_argument]
func ParseMessage(line string) *Message {
	var (
		prefix    string
		command   string
		arguments []string
		final     string
	)

	if pos := strings.Index(line, " :"); pos != -1 {
		arguments, final = strings.Split(line[:pos], " "), line[pos+2:len(line)-2]
	} else {
		arguments = strings.Split(line, " ")
	}

	if arguments[0][0] == ':' && len(arguments) >= 2 {
		prefix = arguments[0]
		command = arguments[1]
		arguments = arguments[2:]
	} else {
		command = arguments[0]
		arguments = arguments[1:]
	}

	return &Message{prefix, command, arguments, final}
}

// server_or_nick[!user@hostname]
// TODO: "server_or_nick!user@" is valid according to this function. Fix this!
// TODO: Maybe also extract user rank (~&@%+)?
func ParsePrefix(source string) (server_or_nick, user, host string) {
	if n := strings.IndexRune(source, '!'); n != -1 {
		server_or_nick = source[:n]
		if m := strings.IndexRune(source, '@'); m != -1 {
			user = source[n+1 : m]
			host = source[m+1:]
		}
	} else {
		server_or_nick = source
	}

	return
}
