package main

import (
	"fmt"
	"os"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/logger"
	nav "github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"
	"github.com/HicaroD/hypersomnia/popup"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var EXIT_MESSAGE string = "an unexpected error ocurred, please see the log file, go to https://github.com/HicaroD/hypersomnia and report an issue!"

var (
	DEFAULT_HTTP_CLIENT_TIMEOUT = 30
)

func exitAppWithUnexpectedError() {
	fmt.Println(EXIT_MESSAGE)
	os.Exit(1)
}

type Hyper struct {
	app          *tview.Application
	navigator    *nav.Navigator
	pageManager  *pages.Manager
	popupManager *popup.Manager
}

func NewHyper(app *tview.Application, navigator *nav.Navigator, pageManager *pages.Manager, popupManager *popup.Manager) *Hyper {
	return &Hyper{
		app:          app,
		navigator:    navigator,
		pageManager:  pageManager,
		popupManager: popupManager,
	}
}

func (h *Hyper) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	// NOTE(bug): Ctrl+Backspace causes Ctrl+H to be triggered, which is totally
	// incorrect
	// It is a possible bug in the Tview source code
	pageIndex, ok := nav.KEY_TO_PAGE[pressedKey]
	if ok {
		page, err := h.pageManager.GetPage(pageIndex)
		if err != nil {
			message := fmt.Sprintf("unable to get page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
			h.errorAndLog(message)
			exitAppWithUnexpectedError()
		}

		err = h.navigator.Navigate(page, false)
		if err != nil {
			message := fmt.Sprintf("unable to navigate to page with index %s due to the following error: %s", pages.NAMES[pageIndex], err)
			h.errorAndLog(message)
			exitAppWithUnexpectedError()
		}
		return event
	}

	switch pressedKey {
	case tcell.KeyEsc:
		if h.navigator.CurrentPage == pages.POPUP {
			h.navigator.Pop()
		}
	}

	return event
}

func (h *Hyper) Run() {
	h.app.SetInputCapture(h.InputCapture)

	welcomePage, err := h.pageManager.GetPage(pages.WELCOME)
	if err != nil {
		message := fmt.Sprintf("unable to get welcome page due to the following error:\n%s", err)
		h.errorAndLog(message)
		return
	}

	err = h.navigator.Navigate(welcomePage, true)
	if err != nil {
		message := fmt.Sprintf("unable to navigate to welcome page due to the following error:\n%s", err)
		h.errorAndLog(message)
		return
	}

	if err := h.app.Run(); err != nil {
		message := fmt.Sprintf("unable to execute application due to the following error:\n%s", err)
		h.errorAndLog(message)
		return
	}
}

func (h *Hyper) errorAndLog(message string) {
	logger.Error.Print(message)
	h.popupManager.ShowPopup(popup.POPUP_ERROR, message)
}

// TODO: this function is too big, I need to refactor this
func main() {
	err := logger.InitLogFile()
	if err != nil {
		logger.Error.Printf("unable to init logger: %s\n", err)
		exitAppWithUnexpectedError()
		return
	}
	defer func() {
		err := logger.Close()
		if err != nil {
			logger.Error.Printf("unable to close log file: %s\n", err)
			exitAppWithUnexpectedError()
			return
		}
	}()

	logger.Info.Println("Logger initialized successfuly")

	// TODO: create database in the configuration folder
	database, err := db.New("endpoints.sqlite")
	if err != nil {
		logger.Error.Printf("unable to open SQLite3 database: %s\n", err)
		exitAppWithUnexpectedError()
		return
	}
	defer func() {
		err := database.Close()
		if err != nil {
			logger.Error.Printf("unable to close SQLite3 database: %s\n", err)
			exitAppWithUnexpectedError()
			return
		}
	}()

	logger.Info.Println("Database initialized successfuly")

	client := hyperHttp.New(DEFAULT_HTTP_CLIENT_TIMEOUT)

	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	hyperPages := tview.NewPages()
	app.SetRoot(hyperPages, true)

	navigator := nav.New(hyperPages)
	ppm := popup.New(navigator.ShowPopup)
	pageManager, err := pages.New(client, database, ppm, navigator.Pop)
	if err != nil {
		exitAppWithUnexpectedError()
		return
	}

	hyper := NewHyper(app, navigator, pageManager, ppm)
	hyper.Run()
}
