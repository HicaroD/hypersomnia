package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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
		pressedKey := event.Key()
		pageIndex, ok := KEY_TO_PAGE[pressedKey]
		if ok {
			hyper.navigator.Navigate(pageIndex)
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
