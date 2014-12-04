package pokemon

import (
	"math/rand"
	"strings"
	"time"

	"github.com/FSX/jun/client"
	"github.com/sorcix/irc"
)

// https://en.wikiquote.org/wiki/Pok%C3%A9mon
// http://www.pokemonquotes.com/brock-quotes/
var quotes []string = []string{
	// Pikachu
	"“PIKA!” — Pikachu",
	"“Pika Pika” — Pikachu",
	"“Pika Pi” — Pikachu",
	"“Chuuuuu” — Pikachu",
	"“Pi-pi-pi” — Pikachu",
	"“PIKA-CHU-UUUUUUU!” — Pikachu",
	"“Pi Pikachu!” — Pikachu",

	// Brock
	"“Hey I know! I'll use my trusty frying pan as a drying pan!” — Brock",
	"“How about some prune juice!?” — Brock",
	"“I believe in rock-hard defense and determination!” — Brock",
	"“Her temper could sure use a little evolution.” — Brock",
	"“Don't you worry, Totodile, I'm gonna teach you how to be successful in love, just like me!” — Brock",
	"“When you have lemons, you make lemonade; and when you have rice, you make rice balls.” — Brock",
	"“For me, summer means bathing suits and girls to wear them!” — Brock",
	"“DON'T YOU GET IT?!! If two Butterfree fall in love, their trainers can meet, and THEY CAN FALL IN LOVE, TOO!” — Brock",
	"“AHH! I wasted a donut!” — Brock",
	"“Ash and Misty, sittin' in a tree... Ha ha ha ha ha!” — Brock",
	"“Time to dry up, water girl!” — Brock",
	"“I never forget the face of a pretty girl! This book helps me remember their names.” — Brock",
	"“This is the only song I know.” — Brock",

	// Misty
	"“If I want your opinion I'LL ASK FOR IT!!!” — Misty",
	"“HE'S NOT MY BOYFRIEND!” — Misty",
	"“Well, Ash Ketchum... finally, I know how you feel about me.” — Misty",
	"“Pikachu, you're a Pika-pal!” — Misty",
	"“Well, my name.....My name is Anne Chovie.” — Misty",

	// Ash Ketchum
	"“Choose it or lose it!” — Ash Ketchum",
	"“I'm twice as good as Gary!” — Ash Ketchum",
	"“It was an egg-cident! Get it? 'Egg?'” — Ash Ketchum",
	"“I have my own method of bending spoons. (bends the spoon with his hands) Ha! Muscle over mind!” — Ash Ketchum",
	"“Please don't eat my hat.” — Ash Ketchum",
	"“I'm having a major hat crisis. Could you try to steal Pikachu some other time?” — Ash Ketchum",
	"“Uh, my... name... is... Ketchup! No, wait! My name is really Tom Ato!” — Ash Ketchum",
	"“Let's eat fast so we can eat again!” — Ash Ketchum",
	"“That mini-Misty is even more scary than she is.” — Ash Ketchum",
	"“You look like a guy anyway!” — Ash Ketchum",
	"“Please don't stare at me like that! I'm a very shy little girl!” — Ash Ketchum",
	"“You guys go, I'll be fine. Just don't bring my mom home too late.” — Ash Ketchum",
	"“Yo, Brocko!” — Ash Ketchum",
	"“I choose you pikachu!” — Ash Ketchum",

	// Prof. Oak
	"“So, tell me about yourself. Are you a boy or a girl?” — Prof. Oak",
	"“This is my grandson. He's been your rival since you were both babies. Err... what was his name again...?” — Prof. Oak",
	"“You look more like you're ready for bed than Pokémon training.” — Prof. Oak",
	"“There's an ongoing debate in the academic community as to whether these Pidgey represent evolution, devolution, or some mutated strain.” — Prof. Oak",
}

func PokemonQuotes(j *client.Client) {
	rand.Seed(time.Now().UnixNano())
	l := len(quotes)

	j.AddCallback("PRIVMSG", func(message *irc.Message) {
		if strings.HasPrefix(message.Trailing, "!pokemon") {
			j.Privmsg(message.Params[0], quotes[rand.Intn(l)])
		}
	})
}
