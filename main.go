package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BACKGROUND_COLOR tcell.Color = tcell.NewRGBColor(28, 28, 28)

//go:embed ascii_art.txt
var WELCOME_MESSAGE string

type Hyper struct {
	app       *tview.Application
	pages     *tview.Pages
	navigator *HyperNavigator
}

func NewHyper() *Hyper {
	app := tview.NewApplication()
	pages := tview.NewPages()
	return &Hyper{app: app, pages: pages, navigator: SetupPages(pages)}
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlP:
			hyper.navigator.Navigate(COLLECTIONS)
			return event
		}
		return event
	})

	hyper.navigator.Navigate(WELCOME)
	hyper.app.SetRoot(hyper.pages, true)
	if err := hyper.app.Run(); err != nil {
		panic(err)
	}
}

func main() {
	app := NewHyper()
	app.Run()
}
