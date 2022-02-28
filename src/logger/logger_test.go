package logger

import "testing"

func TestLog1(t *testing.T) {
	logger := New()
	logger.Info("Grissom")
	logger.Info("Grissom %s", "Hello")
}
