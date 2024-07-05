package pages

import (
	"log"

	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/models"
	"github.com/HicaroD/hypersomnia/popup"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (
	OnRequestCallback       func(hyperHttp.Request) (*hyperHttp.Response, error)
	OnListEndpointsCallback func() ([]*models.Endpoint, error)
	ShowPopupCallback       func(kind popup.PopupKind, text string)
)

type EndpointsPage struct {
	main tview.Primitive

	methods              *tview.DropDown
	url                  *tview.InputField
	body, query, headers *tview.TextArea
	response             *tview.TextView

	endpoints []*models.Endpoint

	onRequest       OnRequestCallback
	onListEndpoints OnListEndpointsCallback
	showPopup       ShowPopupCallback
}

func (page *EndpointsPage) Setup() {
	endpoints, err := page.onListEndpoints()
	if err != nil {
		log.Fatal(err)
	}
	page.endpoints = endpoints

	main := tview.NewFlex()
	main.SetBorder(true)
	main.SetDirection(tview.FlexColumn)
	main.SetBackgroundColor(utils.COLOR_DARK_GREY)

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

			request := hyperHttp.Request{
				Method:      selectedMethod,
				Url:         endpointUrl,
				Body:        body,
				QueryParams: query,
				Headers:     headers,
			}
			response, err := page.onRequest(request)
			if err != nil {
				page.showPopup(popup.POPUP_ERROR, err.Error())
				break
			}

			page.response.SetText(response.Body)
		}
		return event
	})

	requestSection := page.buildRequestSection()
	endpointsSection := page.buildEndpointsSection()
	responseSection := page.buildResponseSection()

	main.AddItem(
		endpointsSection,
		0,
		2,
		false,
	)
	main.AddItem(
		requestSection,
		0,
		4,
		false,
	)
	main.AddItem(
		responseSection,
		0,
		4,
		false,
	)

	page.main = main
}

func (page *EndpointsPage) Index() Index          { return ENDPOINTS }
func (page *EndpointsPage) Page() tview.Primitive { return page.main }

func (page *EndpointsPage) buildEndpointsSection() tview.Primitive {
	endpoints := tview.NewFlex()
	endpoints.SetTitle("Endpoints")
	endpoints.SetBorder(true)
	endpoints.SetBackgroundColor(utils.COLOR_DARK_GREY)
	endpoints.SetDirection(tview.FlexColumn)

	list := tview.NewList()
	list.SetBackgroundColor(utils.COLOR_DARK_GREY)
	list.SetMainTextStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY))
	list.SetSelectedStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY).Bold(true))
	list.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		if index > len(page.endpoints) {
			return
		}
		selectedEndpoint := page.endpoints[index]
		page.setCurrentEndpoint(selectedEndpoint)
	})

	for _, endpointItem := range page.endpoints {
		list.AddItem(endpointItem.String(), "", 0, nil)
	}

	endpoints.AddItem(
		list,
		0,
		1,
		false,
	)
	return endpoints
}

func (page *EndpointsPage) buildRequestSection() tview.Primitive {
	defaultOption := 0 // GET
	methodDropdown := widgets.Dropdown(models.ENDPOINT_METHODS, defaultOption, nil)

	// TODO: set paste handler callback (validator) to only accept links (if necessary)
	urlInput := widgets.InputField("https://google.com/")
	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdown, 0, 1, false).
		AddItem(urlInput, 0, 5, false)

	requestBodyArea := widgets.TextArea("Body")
	queryParametersArea := widgets.TextArea("Query parameters")
	headersArea := widgets.TextArea("Headers")

	requestForm := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(urlForm, 0, 1, false).
		AddItem(queryParametersArea, 0, 2, false).
		AddItem(requestBodyArea, 0, 7, false).
		AddItem(headersArea, 0, 2, false)
	requestForm.SetBorder(true)
	requestForm.SetBackgroundColor(utils.COLOR_DARK_GREY)
	requestForm.SetTitle("Request")

	page.methods = methodDropdown
	page.url = urlInput
	page.body = requestBodyArea
	page.query = queryParametersArea
	page.headers = headersArea

	return requestForm
}

func (page *EndpointsPage) buildResponseSection() tview.Primitive {
	response := tview.NewTextView()
	response.SetTitle("Response")
	response.SetBorder(true)
	response.SetBackgroundColor(utils.COLOR_DARK_GREY)
	page.response = response
	return response
}

func (page *EndpointsPage) setCurrentEndpoint(endpoint *models.Endpoint) {
	page.methods.SetCurrentOption(endpoint.MethodIndex())
	page.url.SetText(endpoint.Url)
	page.query.SetText(endpoint.RequestQueryParams, true)
	page.body.SetText(endpoint.RequestBody, true)
	page.headers.SetText(endpoint.RequestHeaders, true)
}
