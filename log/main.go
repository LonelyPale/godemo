package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Fatalf("failed to server: %v", errors.New("test"))
	log.Fatalf("failed to server: %v", nil)
}
