package rule

type Deck struct {
	seq  []Card
	next int
}

func New() *Deck {
	return new(Deck)
}

func (deck *Deck) Shuffle() {
}
