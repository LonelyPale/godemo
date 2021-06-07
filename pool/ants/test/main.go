package main

import (
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	pool, err := ants.NewPoolWithFunc(3, func(i interface{}) {
		log.Println("run:", i)
		time.Sleep(time.Second)

	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		if i == 9 {
			pool.Release()
		}

		if err := pool.Invoke(i); err != nil {
			log.Println(err)
		}
	}

	time.Sleep(5 * time.Second)
}
