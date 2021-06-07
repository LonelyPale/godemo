package main

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Red     = color.New(color.FgRed)
	Blue    = color.New(color.FgBlue)
	Green   = color.New(color.FgGreen)
	Yellow  = color.New(color.FgYellow)
	Magenta = color.New(color.FgMagenta)
	Cyan    = color.New(color.FgCyan)
	White   = color.New(color.FgWhite)
	Black   = color.New(color.FgBlack)
)

func main() {
	Print(Red, "red\n")
	Print(Blue, "blue\n")
	Print(Green, "green\n")
}

func Print(c *color.Color, s string) {
	if len(s) == 0 {
		return
	}
	if _, err := c.Print(s); err != nil {
		fmt.Println(err)
	}
}
