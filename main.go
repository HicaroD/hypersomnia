package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	nav "github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	logFile     *os.File
	app         *tview.Application
	navigator   *nav.Navigator
	pageManager *pages.Manager
}

func NewHyper(app *tview.Application, navigator *nav.Navigator, pageManager *pages.Manager, logFile *os.File) *Hyper {
	return &Hyper{
		logFile:     logFile,
		app:         app,
		navigator:   navigator,
		pageManager: pageManager,
	}
}

func (hyper *Hyper) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	pageIndex, ok := nav.KEY_TO_PAGE[pressedKey]
	if ok {
		page, err := hyper.pageManager.GetPage(pageIndex)
		if err != nil {
			log.Fatalf("unable to get page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
		}

		err = hyper.navigator.Navigate(page)
		if err != nil {
			log.Fatalf("unable to navigate to page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
		}
		return event
	}

	switch pressedKey {
	case tcell.KeyEsc:
		if hyper.navigator.CurrentPage == pages.POPUP {
			hyper.navigator.Pop()
		}
	}

	return event
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(hyper.InputCapture)

	welcomePage, err := hyper.pageManager.GetPage(pages.WELCOME)
	if err != nil {
		log.Fatalf("unable to get welcome page due to the following error:\n%s", err)
	}

	err = hyper.navigator.Navigate(welcomePage)
	if err != nil {
		log.Fatalf("unable to navigate to welcome page due to the following error:\n%s", err)
	}

	if err := hyper.app.Run(); err != nil {
		log.Fatalf("unable to execute application due to the following error:\n%s", err)
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
	// TODO: passing this file around is boring, is there a way to make it
	// global, so I can log anything anywhere I want?
	logFile, err := buildLogFile()
	if err != nil {
		log.Fatalf("unable to build log file: %s\n", err)
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.Fatalf("unable to close log file: %s\n", err)
		}
	}()

	// TODO: create database in the configuration folder
	database, err := db.New("endpoints.sqlite", logFile)
	if err != nil {
		log.Fatalf("unable to open SQLite3 database: %s\n", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatalf("unable to close SQLite3 database: %s\n", err)
		}
	}()

	client := hyperHttp.New(
		&http.Client{
			// TODO: 30 seconds by default, but user should be able to decide the
			// timeout
			Timeout: 30 * time.Second,
		},
	)

	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	hyperPages := tview.NewPages()
	app.SetRoot(hyperPages, true)

	navigator := nav.New(hyperPages)
	pageManager := pages.New(client, database, navigator.ShowPopup)

	hyper := NewHyper(app, navigator, pageManager, logFile)
	hyper.Run()
}
