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

var PAGE_NAMES map[HyperPageIndex]string = map[HyperPageIndex]string{
	WELCOME:   "welcome",
	ENDPOINTS: "endpoints",
	POPUP:     "popup",
	HELP:      "help",
}

type HyperNavigator struct {
	pages        *tview.Pages
	mapper       map[HyperPageIndex]HyperPage
	currentPage  HyperPageIndex
	previousPage HyperPageIndex
}

func NewNavigator(pages *tview.Pages, db *HyperDB) *HyperNavigator {
	navigator := HyperNavigator{
		pages:        pages,
		mapper:       map[HyperPageIndex]HyperPage{},
		currentPage:  -1,
		previousPage: -1,
	}

	navigator.mapper[WELCOME] = &WelcomePage{}
	navigator.mapper[ENDPOINTS] = &EndpointsPage{navigator: &navigator, db: db}
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

func (navigator *HyperNavigator) Pop() {
	navigator.pages.RemovePage(PAGE_NAMES[navigator.currentPage])
}

func (navigator *HyperNavigator) ShowPopup(popup *tview.Flex) {
	navigator.pages = navigator.pages.AddPage("popup", popup, true, true)
	navigator.previousPage = navigator.currentPage
	navigator.currentPage = POPUP
}
