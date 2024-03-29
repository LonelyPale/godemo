package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const col = 30
	// Clear the screen by printing \x0c.
	//bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	bar := fmt.Sprintf("\x0d[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!\n", strings.Repeat("=", col))
}
