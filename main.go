package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BACKGROUND_COLOR tcell.Color = tcell.NewRGBColor(28, 28, 28)
var WELCOME_MESSAGE string = `
 _   _                                                 _       
| | | |                                               (_)      
| |_| |_   _ _ __   ___ _ __ ___  ___  _ __ ___  _ __  _  __ _ 
|  _  | | | | '_ \ / _ \ '__/ __|/ _ \| '_  _  \| '_ \| |/  _ |
| | | | |_| | |_) |  __/ |  \__ \ (_) | | | | | | | | | | (_| |
\_| |_/\__, | .__/ \___|_|  |___/\___/|_| |_| |_|_| |_|_|\__,_|
        __/ | |                                                
       |___/|_|                                                

  It may seem absurd, but Hypersomnia is better than Insomnia
`

func main() {
	app := tview.NewApplication()

	pages := tview.NewPages()
	pages.SetBackgroundColor(BACKGROUND_COLOR)

	welcomeText :=
		tview.NewTextView().
			SetSize(0, 0).
			SetText(WELCOME_MESSAGE).
			SetTextColor(tcell.ColorWhite).
			SetTextAlign(tview.AlignCenter).
			SetBackgroundColor(BACKGROUND_COLOR).
			SetBorder(true)

	pages.AddPage("", welcomeText, false, true)
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
