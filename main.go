package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/logger"
	nav "github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var EXIT_MESSAGE string = "an unexpected error ocurred, please go to https://github.com/HicaroD/hypersomnia and report an issue!"

func exitAppWithUnexpectedError() {
	fmt.Println(EXIT_MESSAGE)
	os.Exit(1)
}

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
			// TODO: show popup here
			logger.Error.Printf("unable to get page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
			exitAppWithUnexpectedError()
		}

		err = hyper.navigator.Navigate(page)
		if err != nil {
			// TODO: show popup here
			logger.Error.Printf("unable to navigate to page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
			exitAppWithUnexpectedError()
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
		// TODO: show popup here
		logger.Error.Printf("unable to get welcome page due to the following error:\n%s", err)
		return
	}

	err = hyper.navigator.Navigate(welcomePage)
	if err != nil {
		// TODO: show popup here
		logger.Error.Printf("unable to navigate to welcome page due to the following error:\n%s", err)
		return
	}

	if err := hyper.app.Run(); err != nil {
		// TODO: show popup here
		logger.Error.Printf("unable to execute application due to the following error:\n%s", err)
		return
	}
}

func main() {
	err := logger.InitLogFile()
	if err != nil {
		logger.Error.Printf("unable to init logger: %s\n", err)
		return
	}
	defer func() {
		err := logger.Close()
		if err != nil {
			logger.Error.Printf("unable to close log file: %s\n", err)
			return
		}
	}()

	logger.Info.Println("Logger initialized successfuly")

	// TODO: create database in the configuration folder
	database, err := db.New("endpoints.sqlite")
	if err != nil {
		logger.Error.Printf("unable to open SQLite3 database: %s\n", err)
		return
	}
	defer func() {
		err := database.Close()
		if err != nil {
			logger.Error.Printf("unable to close SQLite3 database: %s\n", err)
			return
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
	pageManager, err := pages.New(client, database, navigator.ShowPopup)
	if err != nil {
		return
	}

	hyper := NewHyper(app, navigator, pageManager)
	hyper.Run()
}
