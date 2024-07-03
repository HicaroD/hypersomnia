package pages

import (
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/models"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type OnRequestCallback func(hyperHttp.Request) (*hyperHttp.Response, error)

type EndpointsPage struct {
	Main tview.Primitive

	methods              *tview.DropDown
	url                  *tview.InputField
	body, query, headers *tview.TextArea
	response             *tview.TextView

	onRequest OnRequestCallback
}

func (page *EndpointsPage) Setup() {
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
				Method: selectedMethod,
				Url: endpointUrl,
				Body: body,
				QueryParams: query,
				Headers: headers,
			}
			response, err := page.onRequest(request)
			if err != nil {
				// TODO: show popup
				panic(err)
			}

			page.response.SetText(response.Body)
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

	page.Main = main
}

func (page *EndpointsPage) buildEndpointsSection() tview.Primitive {
	endpoints := tview.NewFlex()
	endpoints.SetBorder(true)
	endpoints.SetDirection(tview.FlexColumn)

	list := tview.NewList()
	list.SetBackgroundColor(utils.COLOR_DARK_GREY)
	list.SetMainTextStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY))
	list.SetSelectedStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY).Bold(true))

	// TODO: implement callback for getting list of endpoints

	// storedEndpoints, err := page.db.ListEndpoints()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	storedEndpoints := make([]*models.Endpoint, 0)

	for _, endpointItem := range storedEndpoints {
		list.AddItem(endpointItem.String(), "", 0, nil)
	}

	endpoints.AddItem(
		list,
		0,
		1,
		false,
	)
	endpoints.SetBackgroundColor(utils.COLOR_DARK_GREY)
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
	methodDropdown := widgets.Dropdown(methods, defaultOption, nil)
	page.methods = methodDropdown

	// TODO: set paste handler callback (validator) to only accept links (if necessary)
	urlInput := widgets.InputField("https://google.com/")
	page.url = urlInput
	urlForm := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(methodDropdown, 0, 1, false).
		AddItem(urlInput, 0, 5, false)

	requestBodyArea := widgets.TextArea("Body")
	page.body = requestBodyArea

	queryParametersArea := widgets.TextArea("Query parameters")
	page.query = queryParametersArea

	headersArea := widgets.TextArea("Headers")
	page.headers = headersArea

	requestForm := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(urlForm, 0, 1, false).
		AddItem(queryParametersArea, 0, 2, false).
		AddItem(requestBodyArea, 0, 7, false).
		AddItem(headersArea, 0, 2, false)
	requestForm.SetBorder(true)
	requestForm.SetBackgroundColor(utils.COLOR_DARK_GREY)
	requestForm.SetTitle("Request")

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

func (page *EndpointsPage) Page() tview.Primitive { return page.Main }
