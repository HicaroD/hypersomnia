package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HyperPage interface {
	Build() (string, tview.Primitive)
}

type WelcomePage struct{}

func (page *WelcomePage) Build() (string, tview.Primitive) {
	welcomeText := tview.NewTextView()
	welcomeText.SetBorder(true)
	welcomeText.SetText(WELCOME_MESSAGE)
	welcomeText.SetTextColor(tcell.ColorDodgerBlue)
	welcomeText.SetTextAlign(tview.AlignCenter)
	welcomeText.SetBackgroundColor(WELCOME_DARK_BACKGROUND)
	return "welcome", welcomeText
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
	methodDropdownInput := tview.NewForm()
	methodDropdownInput.AddDropDown(
		"Method",
		[]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "CONNECT", "OPTIONS", "TRACE"},
		0,
		nil,
	)
	methodDropdownInput.SetBackgroundColor(DARK_GREY)

	urlInput := tview.NewForm()
	urlInput.AddInputField("URL", "", 0, nil, nil)
	urlInput.SetBackgroundColor(DARK_GREY)

	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdownInput, 0, 1, false).
		AddItem(urlInput, 0, 3, false)

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
		AddItem(requestBodyArea, 0, 3, false).
		AddItem(queryParametersArea, 0, 1, false).
		AddItem(headersArea, 0, 1, false)

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
