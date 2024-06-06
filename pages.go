package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"time"

	"net/http"

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

type EndpointsPage struct {
	client *http.Client

	// Widgets
	methods              *tview.DropDown
	url                  *tview.InputField
	body, query, headers *tview.TextArea
	response             *tview.TextView
}

func (page *EndpointsPage) Setup() {
	page.client = &http.Client{
		// Let the user decide the timeout
		Timeout: 30 * time.Second,
	}
}

func (page *EndpointsPage) Build() (string, tview.Primitive) {
	page.Setup()

	main := tview.NewFlex()
	main.SetBorder(true)
	main.SetDirection(tview.FlexColumn)
	main.SetBackgroundColor(DARK_GREY)

	main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlSpace:
			// TODO: get data from form fields
			_, selectedMethod := page.methods.GetCurrentOption()
			url := page.url.GetText()
			// body := page.body.GetText()
			// query := page.query.GetText()
			// headers := page.headers.GetText()

			request, err := http.NewRequest(selectedMethod, url, nil)
			if err != nil {
				panic(err)
			}

			resp, err := page.client.Do(request)
			if err != nil {
				panic(err)
			}

			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			formattedJsonBuffer := &bytes.Buffer{}
			if err := json.Indent(formattedJsonBuffer, respBytes, "", "  "); err != nil {
				panic(err)
			}
			page.response.SetText(formattedJsonBuffer.String())

			// TODO: deal with headers and query parameters
		}
		return event
	})

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
	list.SetMainTextStyle(tcell.StyleDefault.Background(DARK_GREY))
	list.SetShortcutStyle(tcell.StyleDefault.Background(DARK_GREY))
	list.AddItem("First endpoint", "", '0', nil)
	list.AddItem("Second endpoint", "", '1', nil)

	endpoints.AddItem(
		list,
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
	page.methods = methodDropdown

	urlInput := tview.NewInputField()
	urlInput.SetBorder(true)
	urlInput.SetFieldBackgroundColor(DARK_GREY)
	urlInput.SetBackgroundColor(DARK_GREY)
	urlInput.SetPlaceholder("https://google.com/")
	urlInput.SetPlaceholderStyle(tcell.StyleDefault.Background(DARK_GREY))
	urlInput.SetPlaceholderTextColor(tcell.ColorGrey)
	// TODO: set paste handler callback to only accept links (if necessary)
	page.url = urlInput

	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdown, 0, 1, false).
		AddItem(urlInput, 0, 5, false)

	requestBodyArea := tview.NewTextArea()
	requestBodyArea.SetBorder(true)
	requestBodyArea.SetTitle("Body")
	requestBodyArea.SetBackgroundColor(DARK_GREY)
	requestBodyArea.SetTextStyle(tcell.StyleDefault.Background(DARK_GREY))
	page.body = requestBodyArea

	queryParametersArea := tview.NewTextArea()
	queryParametersArea.SetBorder(true)
	queryParametersArea.SetTitle("Query parameters")
	queryParametersArea.SetBackgroundColor(DARK_GREY)
	queryParametersArea.SetTextStyle(tcell.StyleDefault.Background(DARK_GREY))
	page.query = queryParametersArea

	headersArea := tview.NewTextArea()
	headersArea.SetBorder(true)
	headersArea.SetTitle("Headers")
	headersArea.SetBackgroundColor(DARK_GREY)
	headersArea.SetTextStyle(tcell.StyleDefault.Background(DARK_GREY))
	page.headers = headersArea

	requestForm := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(urlForm, 0, 1, false).
		AddItem(requestBodyArea, 0, 7, false).
		AddItem(queryParametersArea, 0, 2, false).
		AddItem(headersArea, 0, 2, false)

	requestForm.SetBorder(true)
	requestForm.SetBackgroundColor(DARK_GREY)
	requestForm.SetTitle("Request")

	return requestForm
}

func (page *EndpointsPage) buildResponseSection() tview.Primitive {
	response := tview.NewTextView()
	response.SetTitle("Response")
	response.SetBorder(true)
	response.SetBackgroundColor(DARK_GREY)
	page.response = response
	return response
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
