package main

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
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
	app := cview.NewApplication()

	welcomeText := cview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(cview.AlignCenter)
	welcomeText.SetBackgroundColor(BACKGROUND_COLOR)

	app.SetRoot(welcomeText, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
