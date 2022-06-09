package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	log.Print("Print log...")
	log.WithField("user", &User{
		Name: "demo",
		Age:  99,
	}).Info("Info log...")

	log.WithField("err", errors.New("test-error"))
	log.WithField("err", errors.New("test-error")).Info("info...")
	log.WithField("err", errors.New("test-error")).Error("failed to http server[Stop]")
	log.WithField("err", errors.New("test-error")).Panic("failed to http server[Start]")

	log.Fatalf("failed to server: %v", errors.New("test"))
	log.Fatalf("failed to server: %v", nil)
}
