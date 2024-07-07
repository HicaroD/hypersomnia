package main

import (
	"log"
	"net/http"
	"time"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/logger"
	nav "github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	app         *tview.Application
	navigator   *nav.Navigator
	pageManager *pages.Manager
}

func NewHyper(app *tview.Application, navigator *nav.Navigator, pageManager *pages.Manager) *Hyper {
	return &Hyper{
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

func main() {
	err := logger.InitLogFile()
	if err != nil {
		log.Fatalf("unable to init logger: %s\n", err)
	}
	defer func() {
		err := logger.Close()
		if err != nil {
			log.Fatalf("unable to close log file: %s\n", err)
		}
	}()

	logger.Info.Println("Logger initialized successfuly")

	// TODO: create database in the configuration folder
	database, err := db.New("endpoints.sqlite")
	if err != nil {
		log.Fatalf("unable to open SQLite3 database: %s\n", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatalf("unable to close SQLite3 database: %s\n", err)
		}
	}()

	logger.Info.Println("Database initialized successfuly")

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

	hyper := NewHyper(app, navigator, pageManager)
	hyper.Run()
}
