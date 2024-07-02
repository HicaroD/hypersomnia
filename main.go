package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HicaroD/hypersomnia/navigator"
	"github.com/HicaroD/hypersomnia/pages"
	db "github.com/HicaroD/hypersomnia/database"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Hyper struct {
	logFile   *os.File
	app       *tview.Application
	navigator *navigator.Navigator
}

func NewHyper(db *db.Database, logFile *os.File) *Hyper {
	app := tview.NewApplication()
	app.EnablePaste(true)
	app.EnableMouse(true)

	pages := tview.NewPages()
	app.SetRoot(pages, true)

	return &Hyper{logFile: logFile, app: app, navigator: navigator.New(pages, db)}
}

func (hyper *Hyper) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	pressedKey := event.Key()
	pageIndex, ok := navigator.KEY_TO_PAGE[pressedKey]
	if ok {
		hyper.navigator.Navigate(pageIndex)
	}
	return event
}

func (hyper *Hyper) Run() {
	hyper.app.SetInputCapture(hyper.InputCapture)
	hyper.navigator.Navigate(pages.WELCOME)
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
	// TODO: passing this file around is boring, is there a way to make it
	// global, so I can log anything at any place
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

	app := NewHyper(db, logFile)
	app.Run()
}
