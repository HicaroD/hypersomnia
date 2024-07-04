package pages

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/HicaroD/hypersomnia/utils"
)

//go:embed help.txt
var HELP_MESSAGE string

type HelpPage struct{
	main *tview.TextView
}

func (page *HelpPage) Setup() {
	help := tview.NewTextView()
	help.SetBorder(true)
	help.SetText(HELP_MESSAGE)
	help.SetTextColor(tcell.ColorDodgerBlue)
	help.SetTextAlign(tview.AlignCenter)
	help.SetBackgroundColor(utils.COLOR_WELCOME_DARK_BACKGROUND)
	page.main = help
}

func (page *HelpPage) Page() tview.Primitive { return page.main }
