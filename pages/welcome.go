package pages

import (
	_ "embed"

	"github.com/HicaroD/hypersomnia/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed ascii_art.txt
var WELCOME_MESSAGE string

type WelcomePage struct {
	main *tview.TextView
}

func (page *WelcomePage) Setup() error {
	welcomeText := tview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(tview.AlignCenter)
	welcomeText.SetBackgroundColor(utils.COLOR_WELCOME_DARK_BACKGROUND)

	page.main = welcomeText
	return nil
}

func (page *WelcomePage) Index() Index          { return WELCOME }
func (page *WelcomePage) Page() tview.Primitive { return page.main }
