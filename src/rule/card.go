package rule

import (
	"fmt"
	"strconv"
)

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

func TypeStr(cardType CardType) string {
	switch cardType {
	case Spade:
		return "S"
	case Heart:
		return "H"
	case Club:
		return "C"
	case Diamond:
		return "D"
	}
	return "X"
}

func NumStr(number CardNumber) string {
	if number < 8 {
		return strconv.Itoa(int(number + 2))
	} else {
		switch number {
		case Ten:
			return "T"
		case Jack:
			return "J"
		case Queen:
			return "Q"
		case King:
			return "K"
		case Ace:
			return "A"
		}
	}
	return "0"
}

func (card *Card) Show() string {
	outputStr := fmt.Sprintf("%s%s", TypeStr(card.Type), NumStr(card.No))
	log.Info(outputStr)
	return outputStr
}
