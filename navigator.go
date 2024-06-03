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
)

type HyperPage struct {
	name string
	page tview.Primitive
}

type HyperNavigator struct {
	pages       *tview.Pages
	mapper      map[HyperPageIndex]HyperPage
	currentPage HyperPageIndex
}

func SetupPages(pages *tview.Pages) *HyperNavigator {
	hyperPages := HyperNavigator{
		pages:       pages,
		mapper:      map[HyperPageIndex]HyperPage{},
		currentPage: -1,
	}

	hyperPages.mapper[WELCOME] = buildWelcomePage()
	hyperPages.mapper[ENDPOINTS] = buildEndpointsPage()

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
	navigator.currentPage = index
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

func buildEndpointsPage() HyperPage {
	collections := tview.NewFlex()
	collections.SetBorder(true)
	collections.SetDirection(tview.FlexColumn)
	collections.SetBackgroundColor(BACKGROUND_COLOR)

	endpoints := tview.NewFlex()
	endpoints.SetBorder(true)
	endpoints.SetDirection(tview.FlexColumn)
	endpoints.SetBackgroundColor(BACKGROUND_COLOR)
	endpoints.AddItem(
		tview.NewList().
			AddItem("First endpoint", "", '0', nil),
		0,
		1,
		false,
	)
	collections.AddItem(endpoints, 0, 2, false)

	collections.AddItem(
		tview.NewBox().
			SetTitle("Request").
			SetBorder(true).
			SetBackgroundColor(BACKGROUND_COLOR),
		0,
		4,
		false,
	)

	collections.AddItem(tview.NewBox().SetTitle("Response").SetBorder(true), 0, 4, false)

	return HyperPage{"collections", collections}
}
