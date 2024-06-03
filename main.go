package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BACKGROUND_COLOR tcell.Color = tcell.NewRGBColor(28, 28, 28)

//go:embed ascii_art.txt
var WELCOME_MESSAGE string

func main() {
	app := tview.NewApplication()

	welcomeText := tview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(tview.AlignCenter)
	welcomeText.SetBackgroundColor(BACKGROUND_COLOR)

	app.SetRoot(welcomeText, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
