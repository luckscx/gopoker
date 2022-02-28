package rule

type CardType int
type CardNumber int

const (
	Spade CardType = iota
	Heart
	Diamond
	Club
)

const (
	Two CardNumber = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Type CardType
	No   CardNumber
}
