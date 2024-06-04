package main

import (
	_ "embed"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed ascii_art.txt
var WELCOME_MESSAGE string

type Hyper struct {
	app       *tview.Application
	navigator *HyperNavigator
}

func NewHyper() *Hyper {
	app := tview.NewApplication()
	pages := tview.NewPages()
	app.SetRoot(pages, true)
	return &Hyper{app: app, navigator: SetupPages(pages)}
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlE:
			hyper.navigator.Navigate(ENDPOINTS)
			return event
		}
		return event
	})

	hyper.navigator.Navigate(WELCOME)

	if err := hyper.app.EnablePaste(true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func main() {
	app := NewHyper()
	app.Run()
}
