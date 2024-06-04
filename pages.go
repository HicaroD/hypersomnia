package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HyperPage interface {
	Build() (string, tview.Primitive)
}

type WelcomePage struct{}

func (page WelcomePage) Build() (string, tview.Primitive) {
	welcomeText := tview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(tview.AlignCenter)
	welcomeText.SetBackgroundColor(BACKGROUND_COLOR)
	return "welcome", welcomeText
}

type EndpointsPage struct{}

func (page EndpointsPage) Build() (string, tview.Primitive) {
	collections := tview.NewFlex()
	collections.SetBorder(true)
	collections.SetDirection(tview.FlexColumn)
	collections.SetBackgroundColor(BACKGROUND_COLOR)

	endpoints := tview.NewFlex()
	endpoints.SetBorder(true)
	endpoints.SetDirection(tview.FlexColumn)
	endpoints.SetBackgroundColor(BACKGROUND_COLOR)
	endpoints.AddItem(
		tview.NewList().
			AddItem("First endpoint", "", '0', nil),
		0,
		1,
		false,
	)
	collections.AddItem(endpoints, 0, 2, false)

	collections.AddItem(
		tview.NewBox().
			SetTitle("Request").
			SetBorder(true).
			SetBackgroundColor(BACKGROUND_COLOR),
		0,
		4,
		false,
	)

	collections.AddItem(tview.NewBox().SetTitle("Response").SetBorder(true), 0, 4, false)

	return "collections", collections
}
