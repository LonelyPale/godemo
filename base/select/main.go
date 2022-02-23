package main

import "time"

func main() {
	go func() {
		println("starting")
		time.Sleep(time.Second * 3)
		println("done")
	}()
	select {}
}
