package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"net/http"
	"net/url"

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
	return PAGE_NAMES[WELCOME], welcome
}

type EndpointsPage struct {
	navigator *HyperNavigator

	client *http.Client

	// Widgets
	methods              *tview.DropDown
	url                  *tview.InputField
	body, query, headers *tview.TextArea
	response             *tview.TextView
}

func (page *EndpointsPage) Setup() {
	page.client = &http.Client{
		// TODO: 30 seconds by default, but user should be able to decide the
		// timeout
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
			// TODO: be careful with unnecessary memory usage
			//       If the user start spamming Ctrl+Space, it will keep allocating
			//       strings in memory because of the method "GetText"
			//       Make sure to guarantee I'm only calling this method when
			//       necessary by keeping tracking of the state of each input in
			//       order to verify if any of these inputs has changed since the
			//       last time Ctrl+A was pressed
			_, selectedMethod := page.methods.GetCurrentOption()
			endpointUrl := page.url.GetText()
			body := page.body.GetText()
			headers := page.headers.GetText()
			query := page.query.GetText()

			request, err := http.NewRequest(selectedMethod, endpointUrl, strings.NewReader(body))
			if err != nil {
				page.navigator.ShowPopup(HyperPopup(ERROR, "Unable to build request"))
				break
			}

			err = addQueryParams(request, query)
			// TODO(errors)
			if err != nil {
				page.navigator.ShowPopup(HyperPopup(ERROR, "Invalid format for query parameters"))
				break
			}

			err = addHeaders(request, headers)
			// TODO(errors)
			if err != nil {
				page.navigator.ShowPopup(HyperPopup(ERROR, "Invalid format for headers"))
				break
			}

			resp, err := page.client.Do(request)
			// TODO(errors)
			if err != nil {
				requestErr := err.(*url.Error)
				errorMessage := fmt.Sprintf("Unable to do HTTP request due to %s\n", requestErr.Err)
				page.navigator.ShowPopup(HyperPopup(ERROR, errorMessage))
				break
			}

			// TODO: deal with other kind of responses, not only JSON
			respBytes, err := io.ReadAll(resp.Body)
			// TODO(errors)
			if err != nil {
				page.navigator.ShowPopup(HyperPopup(ERROR, "Unable to read body HTTP request"))
				break
			}

			formattedJsonBuffer := &bytes.Buffer{}
			// TODO(errors)
			err = json.Indent(formattedJsonBuffer, respBytes, "", "  ")
			if err != nil {
				page.navigator.ShowPopup(HyperPopup(ERROR, "Unable to format JSON from response body"))
				break
			}

			page.response.SetText(formattedJsonBuffer.String())
			// TODO: deal with body, headers and query parameters
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

	return PAGE_NAMES[ENDPOINTS], main
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
	methodDropdown := HyperDropdown(methods, defaultOption, nil)
	page.methods = methodDropdown

	// TODO: set paste handler callback (validator) to only accept links (if necessary)
	urlInput := HyperInputField("https://google.com/")
	page.url = urlInput
	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdown, 0, 1, false).
		AddItem(urlInput, 0, 5, false)

	requestBodyArea := HyperTextArea("Body")
	page.body = requestBodyArea

	queryParametersArea := HyperTextArea("Query parameters")
	page.query = queryParametersArea

	headersArea := HyperTextArea("Headers")
	page.headers = headersArea

	requestForm := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(urlForm, 0, 1, false).
		AddItem(queryParametersArea, 0, 2, false).
		AddItem(requestBodyArea, 0, 7, false).
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
	return PAGE_NAMES[HELP], help
}
