package rule

import (
	"logger"
	"math/rand"
	"time"
)

type Deck struct {
	seq  []Card
	next int
}

var log *logger.Logger

func init() {
	log = logger.New()
	rand.Seed(time.Now().UnixNano())
}

func NewDeck() *Deck {
	deck := new(Deck)
	for i := Spade; i <= Club; i++ {
		for j := Two; j <= Ace; j++ {
			card := Card{i, j}
			deck.seq = append(deck.seq, card)
		}
	}
	return deck
}

func (deck *Deck) Print() {
	s := len(deck.seq)
	for i := 0; i < s; i++ {
		card := deck.seq[i]
		card.Show()
	}
}

func (deck *Deck) Shuffle() {
	s := len(deck.seq)
	for i := 0; i < s; i++ {
		r := rand.Intn(s)
		deck.seq[i], deck.seq[r] = deck.seq[r], deck.seq[i]
	}
}
