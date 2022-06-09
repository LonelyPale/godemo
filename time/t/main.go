package main

import (
	"fmt"
	"time"
)

func main() {

	// Declaring t and u in UTC
	t := time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)
	u := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)

	// Calling Before method
	res := t.Before(u)

	// Prints output
	fmt.Printf("%v", res)
}
