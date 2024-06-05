package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed ascii_art.txt
var WELCOME_MESSAGE string

//go:embed help.txt
var HELP_MESSAGE string

type HyperPage interface {
	Build() (string, tview.Primitive)
}

type WelcomePage struct{}

func (page *WelcomePage) Build() (string, tview.Primitive) {
	welcome := tview.NewTextView()
	welcome.SetBorder(true)
	welcome.SetText(WELCOME_MESSAGE)
	welcome.SetTextColor(tcell.ColorDodgerBlue)
	welcome.SetTextAlign(tview.AlignCenter)
	welcome.SetBackgroundColor(WELCOME_DARK_BACKGROUND)
	return "welcome", welcome
}

type EndpointsPage struct{}

func (page *EndpointsPage) Build() (string, tview.Primitive) {
	main := tview.NewFlex()
	main.SetBorder(true)
	main.SetDirection(tview.FlexColumn)
	main.SetBackgroundColor(DARK_GREY)

	endpointsSection := page.buildEndpointsSection()
	main.AddItem(
		endpointsSection,
		0,
		2,
		false,
	)

	requestSection := page.buildRequestSection()
	main.AddItem(
		requestSection,
		0,
		4,
		false,
	)

	responseSection := page.buildResponseSection()
	main.AddItem(
		responseSection,
		0,
		4,
		false,
	)

	return "collections", main
}

func (page *EndpointsPage) buildEndpointsSection() tview.Primitive {
	endpoints := tview.NewFlex()
	endpoints.SetBorder(true)
	endpoints.SetDirection(tview.FlexColumn)
	list := tview.NewList()
	list.SetBackgroundColor(DARK_GREY)
	endpoints.AddItem(
		list.AddItem("First endpoint", "", '0', nil).
			AddItem("Second endpoint", "", '1', nil),
		0,
		1,
		false,
	)
	endpoints.SetBackgroundColor(DARK_GREY)
	endpoints.SetTitle("Endpoints")
	return endpoints
}

func (page *EndpointsPage) buildRequestSection() tview.Primitive {
	methodDropdown := tview.NewDropDown()
	methods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
		"HEAD",
		"CONNECT",
		"OPTIONS",
		"TRACE",
	}
	defaultOption := 0 // GET
	methodDropdown.SetOptions(
		methods,
		// TODO: selected callback
		nil,
	)
	methodDropdown.SetCurrentOption(defaultOption)
	methodDropdown.SetFieldBackgroundColor(DARK_GREY)
	methodDropdown.SetBorder(true)
	methodDropdown.SetBackgroundColor(DARK_GREY)

	urlInput := tview.NewInputField()
	urlInput.SetBorder(true)
	urlInput.SetFieldBackgroundColor(DARK_GREY)
	urlInput.SetBackgroundColor(DARK_GREY)
	urlInput.SetPlaceholder("https://google.com/")
	urlInput.SetPlaceholderStyle(tcell.StyleDefault.Background(DARK_GREY))
	urlInput.SetPlaceholderTextColor(tcell.ColorGrey)
	// TODO: set paste handler callback to only accept links (if necessary)

	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdown, 0, 1, false).
		AddItem(urlInput, 0, 5, false)

	queryParametersArea := tview.NewTextArea()
	queryParametersArea.SetBorder(true)
	queryParametersArea.SetTitle("Query parameters")
	queryParametersArea.SetBackgroundColor(DARK_GREY)

	requestBodyArea := tview.NewTextArea()
	requestBodyArea.SetBorder(true)
	requestBodyArea.SetTitle("Body")
	requestBodyArea.SetBackgroundColor(DARK_GREY)

	headersArea := tview.NewTextArea()
	headersArea.SetBorder(true)
	headersArea.SetTitle("Headers")
	headersArea.SetBackgroundColor(DARK_GREY)

	requestForm := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(urlForm, 0, 1, false).
		AddItem(requestBodyArea, 0, 6, false).
		AddItem(queryParametersArea, 0, 2, false).
		AddItem(headersArea, 0, 2, false)

	requestForm.SetBorder(true)
	requestForm.SetBackgroundColor(DARK_GREY)
	requestForm.SetTitle("Request")

	return requestForm
}

func (page *EndpointsPage) buildResponseSection() tview.Primitive {
	responseBox := tview.NewBox()
	responseBox.SetTitle("Response")
	responseBox.SetBorder(true)
	responseBox.SetBackgroundColor(DARK_GREY)
	return responseBox
}

type HelpPage struct{}

func (page *HelpPage) Build() (string, tview.Primitive) {
	help := tview.NewTextView()
	help.SetBorder(true)
	help.SetText(HELP_MESSAGE)
	help.SetTextColor(tcell.ColorDodgerBlue)
	help.SetTextAlign(tview.AlignCenter)
	help.SetBackgroundColor(WELCOME_DARK_BACKGROUND)
	return "help", help
}
