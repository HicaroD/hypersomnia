package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	logFile   *os.File
	app       *tview.Application
	navigator *navigator.Navigator
}

func NewHyper(pm *pages.Manager, logFile *os.File) *Hyper {
	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	pages := tview.NewPages()
	app.SetRoot(pages, true)

	return &Hyper{logFile: logFile, app: app, navigator: navigator.New(pages, pm)}
}

func (hyper *Hyper) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	pageIndex, ok := navigator.KEY_TO_PAGE[pressedKey]
	if ok {
		err := hyper.navigator.Navigate(pageIndex)
		if err != nil {
			log.Fatalf("unable to navigate to page with index %s due to the following error: %s", pageIndex, err)
		}
	}
	return event
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(hyper.InputCapture)
	err := hyper.navigator.Navigate(pages.WELCOME)
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

func buildPageManager(database *db.Database) *pages.Manager {
	client := hyperHttp.New(
		&http.Client{
			// TODO: 30 seconds by default, but user should be able to decide the
			// timeout
			Timeout: 30 * time.Second,
		},
	)
	return pages.New(client, database)
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
	db, err := db.New("endpoints.sqlite", logFile)
	if err != nil {
		log.Fatalf("unable to open SQLite3 database: %s\n", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("unable to close SQLite3 database: %s\n", err)
		}
	}()

	pm := buildPageManager(db)
	app := NewHyper(pm, logFile)
	app.Run()
}
