package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	app       *tview.Application
	navigator *HyperNavigator
}

func NewHyper() *Hyper {
	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	pages := tview.NewPages()
	app.SetRoot(pages, true)

	return &Hyper{app: app, navigator: NewNavigator(pages)}
}

func (hyper *Hyper) InputHandler(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	pageIndex, ok := KEY_TO_PAGE[pressedKey]
	if ok {
		hyper.navigator.Navigate(pageIndex)
	}

	switch pressedKey {
	case tcell.KeyESC:
		if hyper.navigator.currentPage == POPUP {
			hyper.navigator.Pop()
		}
	}

	return event
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(hyper.InputHandler)
	hyper.navigator.Navigate(WELCOME)
	if err := hyper.app.Run(); err != nil {
		log.Fatalf("Unable to execute application due to the following error:\n%s", err)
	}
}

func main() {
	app := NewHyper()
	app.Run()
}
