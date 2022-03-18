package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	if err := os.Setenv("FYNE_FONT", "/Users/wyb/Library/Fonts/FZZJ-LongYTJF.TTF"); err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(300, 300))

	text := canvas.NewText("Text Object -- 测试", color.Black)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	text.TextSize = 20

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5

	w.SetContent(container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Hello World!"),
		widget.NewLabel("山河辽阔，人间烟火。月光所照，皆是中国！"),
		myText("山河辽阔，"),
		myText("人间烟火。"),
		myText("月光所照，"),
		myText("皆是中国！"),
		layout.NewSpacer(),
		line,
		text,
		widget.NewButton("按钮", func() {
			dialog.ShowInformation("标题-123", "内容-abc", w)
		}),
	))
	w.ShowAndRun()

	if err := os.Unsetenv("FYNE_FONT"); err != nil {
		panic(err)
	}
}

func myText(str string) *canvas.Text {
	text := canvas.NewText(str, color.Black)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	text.TextSize = 25
	return text
}
