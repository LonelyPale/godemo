package main

import (
	"fmt"
	expect "github.com/google/goexpect"
	log "github.com/sirupsen/logrus"
)

func main() {
	cmd := "sudo ls -la"

	e, _, err := expect.Spawn(cmd, -1)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := e.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(e.String())

	buff := make([]byte, 1024)
	n, err := e.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, string(buff))
}
