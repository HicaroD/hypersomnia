package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HyperPageIndex int

const (
	WELCOME HyperPageIndex = iota
	COLLECTIONS
)

type HyperPage struct {
	name string
	page tview.Primitive
}

type HyperNavigator struct {
	pages  *tview.Pages
	mapper map[HyperPageIndex]HyperPage
}

func SetupPages(pages *tview.Pages) *HyperNavigator {
	hyperPages := HyperNavigator{pages: pages, mapper: map[HyperPageIndex]HyperPage{}}

	hyperPages.mapper[WELCOME] = buildWelcomePage()
	hyperPages.mapper[COLLECTIONS] = buildCollectionsPage()

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
	navigator.pages = navigator.pages.AddAndSwitchToPage(page.name, page.page, true)
}

func buildWelcomePage() HyperPage {
	welcomeText := tview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(tview.AlignCenter)
	welcomeText.SetBackgroundColor(BACKGROUND_COLOR)
	return HyperPage{"welcome", welcomeText}
}

func buildCollectionsPage() HyperPage {
	tv := tview.NewTextView().SetText("WASSUP MY BOY")
	return HyperPage{"collections", tv}
}
