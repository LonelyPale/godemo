package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	if err := os.Setenv("FYNE_FONT", "/Users/wyb/Library/Fonts/FZZJ-LongYTJF.TTF"); err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("古诗")
	w.Resize(fyne.NewSize(640, 480))
	w.SetContent(gushiUI())
	w.ShowAndRun()

	if err := os.Unsetenv("FYNE_FONT"); err != nil {
		panic(err)
	}
}

type GuShi struct {
	Title    string
	Author   string
	Dynasty  string
	Contents []string
}

var gushi = GuShi{
	Title:    "出塞",
	Author:   "王昌龄",
	Dynasty:  "〔唐代〕",
	Contents: []string{"秦时明月汉时关，", "万里长征人未还。", "但使龙城飞将在，", "不教胡马度阴山。"},
}

func gushiUI() fyne.CanvasObject {
	c := container.New(layout.NewVBoxLayout())
	rect := canvas.NewRectangle(color.White)
	rect.SetMinSize(fyne.NewSize(1, 50))
	c.Add(rect)
	c.Add(myText(gushi.Title))
	c.Add(myText(gushi.Author + " " + gushi.Dynasty))
	for _, text := range gushi.Contents {
		c.Add(myText(text))
	}
	return c
}

func myText(str string) *canvas.Text {
	text := canvas.NewText(str, color.Black)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	text.TextSize = 38
	return text
}
