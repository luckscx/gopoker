package main

import (
	"logger"
	"rule"
)

var log *logger.Logger

func init() {
	log = logger.New()
}

func main() {
	deck := rule.NewDeck()
	deck.Print()
	log.Info("==================================")
	deck.Shuffle()
	//deck.Print()
}
