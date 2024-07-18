package navigator

import (
	"github.com/HicaroD/hypersomnia/pages"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var KEY_TO_PAGE map[tcell.Key]pages.Index = map[tcell.Key]pages.Index{
	tcell.KeyCtrlW: pages.WELCOME,
	tcell.KeyCtrlE: pages.ENDPOINTS,
	tcell.KeyCtrlH: pages.HELP,
	tcell.KeyCtrlN: pages.NEW_COLLECTION,
}

type Navigator struct {
	pages *tview.Pages

	CurrentPage  pages.Index
	PreviousPage pages.Index
}

func New(p *tview.Pages) *Navigator {
	navigator := Navigator{
		pages:        p,
		CurrentPage:  -1,
		PreviousPage: -1,
	}
	return &navigator
}

func (navigator *Navigator) Navigate(page pages.Page, addAndSwitch bool) error {
	pageIndex := page.Index()
	name := pages.NAMES[pageIndex]
	if addAndSwitch {
		navigator.pages = navigator.pages.AddAndSwitchToPage(name, page.Page(), true)
	} else {
		navigator.pages = navigator.pages.AddPage(name, page.Page(), true, true)
	}
	navigator.PreviousPage = navigator.CurrentPage
	navigator.CurrentPage = pageIndex
	return nil
}

func (navigator *Navigator) Pop() {
	navigator.pages.RemovePage(pages.NAMES[navigator.CurrentPage])
}

func (navigator *Navigator) ShowPopup(popup tview.Primitive) {
	navigator.pages = navigator.pages.AddPage("popup", popup, true, true)
	navigator.PreviousPage = navigator.CurrentPage
	navigator.CurrentPage = pages.POPUP
}
