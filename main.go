package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	logFile   *os.File
	app       *tview.Application
	navigator *HyperNavigator
}

func NewHyper(logFile *os.File) *Hyper {
	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	pages := tview.NewPages()
	app.SetRoot(pages, true)

	return &Hyper{app: app, navigator: NewNavigator(pages), logFile: logFile}
}

func (hyper *Hyper) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	pageIndex, ok := KEY_TO_PAGE[pressedKey]
	if ok {
		hyper.navigator.Navigate(pageIndex)
	}

	switch pressedKey {
	case tcell.KeyESC:
		// TODO: close popup when it loses focus (with mouse interaction),
		// which means I should close the popup when it is clicked outside
		// of the popup
		if hyper.navigator.currentPage == POPUP {
			hyper.navigator.Pop()
		}
	}

	return event
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(hyper.InputCapture)
	hyper.navigator.Navigate(WELCOME)
	if err := hyper.app.Run(); err != nil {
		log.Fatalf("Unable to execute application due to the following error:\n%s", err)
	}
}

func buildLogFile() (*os.File, error) {
	logFile, err := os.Create("log.txt")
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(logFile, "------------------ %s ------------------\n", time.Now().Local())
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func main() {
	// TODO: create log file in the configuration folder
	logFile, err := buildLogFile()
	if err != nil {
		log.Fatalf("unable to build log file: %s\n", err)
	}
	app := NewHyper(logFile)
	app.Run()
}
