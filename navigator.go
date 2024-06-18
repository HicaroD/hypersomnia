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
	pages            *tview.Pages
	mapper           map[HyperPageIndex]HyperPage
	currentPage      HyperPageIndex
	previousPage     HyperPageIndex
}

func NewNavigator(pages *tview.Pages) *HyperNavigator {
	navigator := HyperNavigator{
		pages:            pages,
		mapper:           map[HyperPageIndex]HyperPage{},
		currentPage:      -1,
		previousPage:     -1,
	}

	navigator.mapper[WELCOME] = &WelcomePage{}
	navigator.mapper[ENDPOINTS] = &EndpointsPage{navigator: &navigator}
	navigator.mapper[HELP] = &HelpPage{}

	return &navigator
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

// TODO: Hypersomnia should keep track of the state for all pages
func (navigator *HyperNavigator) Pop() {
	if navigator.previousPage == -1 {
		return
	}
	navigator.Navigate(navigator.previousPage)
}

func (navigator *HyperNavigator) ShowPopup(popup *tview.Flex) {
	navigator.pages = navigator.pages.AddPage("popup", popup, true, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = POPUP
}
