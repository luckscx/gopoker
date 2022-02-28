package main

import (
	"logger"
)

func main() {
	log := logger.New()
	log.Info("Grissom")
	log.Info("Grissom %s", "Hello")
}
