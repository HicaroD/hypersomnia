package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HyperPageIndex int

const (
	WELCOME HyperPageIndex = iota
	ENDPOINTS
	POPUP
	HELP
)

var KEY_TO_PAGE map[tcell.Key]HyperPageIndex = map[tcell.Key]HyperPageIndex{
	tcell.KeyCtrlW: WELCOME,
	tcell.KeyCtrlE: ENDPOINTS,
	tcell.KeyCtrlH: HELP,
}

type HyperNavigator struct {
	pages        *tview.Pages
	mapper       map[HyperPageIndex]HyperPage
	currentPage  HyperPageIndex
	previousPage HyperPageIndex
}

func NewNavigator(pages *tview.Pages) *HyperNavigator {
	hyperPages := HyperNavigator{
		pages:        pages,
		mapper:       map[HyperPageIndex]HyperPage{},
		currentPage:  -1,
		previousPage: -1,
	}

	hyperPages.mapper[WELCOME] = &WelcomePage{}
	hyperPages.mapper[ENDPOINTS] = &EndpointsPage{}
	hyperPages.mapper[HELP] = &HelpPage{}

	return &hyperPages
}

func (navigator *HyperNavigator) GetPage(index HyperPageIndex) HyperPage {
	page, ok := navigator.mapper[index]
	if !ok {
		panic(fmt.Sprintf("invalid index for page: %d\n", index))
	}
	return page
}

func (navigator *HyperNavigator) Navigate(index HyperPageIndex) {
	page := navigator.GetPage(index)
	name, pageWidget := page.Build()
	navigator.pages = navigator.pages.AddAndSwitchToPage(name, pageWidget, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = index
}

func (navigator *HyperNavigator) ShowPopup(popup *tview.Flex) {
	navigator.pages = navigator.pages.AddPage("popup", popup, true, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = POPUP
}

// TODO: if necessary, implement a queue for keeping track of
// pages
func (navigator *HyperNavigator) Pop() {
	if navigator.previousPage == -1 {
		return
	}
	navigator.Navigate(navigator.previousPage)
}
