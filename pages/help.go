package pages

import (
	_ "embed"
	"github.com/HicaroD/hypersomnia/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed help.txt
var HELP_MESSAGE string

type HelpPage struct {
	main *tview.TextView
}

func (page *HelpPage) Setup() error {
	help := tview.NewTextView()
	help.SetBorder(true)
	help.SetText(HELP_MESSAGE)
	help.SetTextColor(tcell.ColorDodgerBlue)
	help.SetTextAlign(tview.AlignCenter)
	help.SetBackgroundColor(utils.COLOR_WELCOME_DARK_BACKGROUND)
	page.main = help
	return nil
}

func (page *HelpPage) Index() Index          { return HELP }
func (page *HelpPage) Page() tview.Primitive { return page.main }
