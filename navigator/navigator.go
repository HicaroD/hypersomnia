package navigator

import (
	"fmt"

	db "github.com/HicaroD/hypersomnia/database"
	"github.com/HicaroD/hypersomnia/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var KEY_TO_PAGE map[tcell.Key]pages.Index = map[tcell.Key]pages.Index{
	tcell.KeyCtrlW: pages.WELCOME,
	tcell.KeyCtrlE: pages.ENDPOINTS,
	tcell.KeyCtrlH: pages.HELP,
}

type Navigator struct {
	pages *tview.Pages

	pageManager  *pages.PageManager
	currentPage  pages.Index
	previousPage pages.Index
}

func New(p *tview.Pages, database *db.Database) *Navigator {
	navigator := Navigator{
		pages:        p,
		pageManager:  pages.New(),
		currentPage:  -1,
		previousPage: -1,
	}
	return &navigator
}

func (navigator *Navigator) Navigate(index pages.Index) error {
	name, page, err := navigator.getPage(index)
	if err != nil {
		return err
	}
	navigator.pages = navigator.pages.AddAndSwitchToPage(name, page, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = index
	return nil
}

func (navigator *Navigator) Pop() {
	navigator.pages.RemovePage(pages.NAMES[navigator.currentPage])
}

func (navigator *Navigator) ShowPopup(popup *tview.Flex) {
	navigator.pages = navigator.pages.AddPage("popup", popup, true, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = pages.POPUP
}

func (navigator *Navigator) getPage(index pages.Index) (string, tview.Primitive, error) {
	var page tview.Primitive
	switch index {
	case pages.WELCOME:
		page = navigator.pageManager.Welcome.Page()
	case pages.HELP:
		page = navigator.pageManager.Help.Page()
	default:
		return "", nil, fmt.Errorf("unimplemented page: %s", index)
	}

	name, ok := pages.NAMES[index]
	if !ok {
		return "", nil, fmt.Errorf("page '%s' name not found", index)
	}

	return name, page, nil
}
