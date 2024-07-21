package pages

import (
	"fmt"

	"github.com/HicaroD/hypersomnia/popup"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewEndpoint struct {
	main       *tview.Flex
	collection *tview.DropDown
	method     *tview.DropDown
	endpoint   *tview.InputField

	onAddNewEndpoint OnAddNewEndpointCallback
	onShowPopup      OnShowPopupCallback
	onPopPage        OnPopPageCallback
}

func (page *NewEndpoint) Setup() error {
	// TODO: get list of collections from database
	// TODO: notify user with popup to add a new collection before adding a new
	// endpoint (in case he does not have one)
	collectionsDropdown := widgets.Dropdown(
		[]string{"my collection placeholder"},
		0,
		func(text string, index int) {
		},
	)

	methodDropdown := widgets.HttpMethodsDropdown()

	endpointUrl := widgets.InputField("https://google.com")
	endpointWidget := widgets.Row([]widgets.Item{
		{
			Item:       methodDropdown,
			Proportion: 1,
		},
		{
			Item:       endpointUrl,
			Proportion: 4,
		},
	})

	buttons := widgets.Row([]widgets.Item{
		{
			Item: widgets.TextButton(
				"Confirm",
				utils.COLOR_DARK_GREY,
				/* border= */ true,
				utils.COLOR_GREEN,
				page.addNewEndpoint,
			),
			Proportion: 1,
		},
		{
			Item: widgets.TextButton(
				"Cancel",
				utils.COLOR_DARK_GREY,
				/* border= */ true,
				utils.COLOR_RED,
				page.onPopPage,
			),
			Proportion: 1,
		},
	})

	items := []widgets.Item{
		{
			Item:       collectionsDropdown,
			Proportion: 1,
		},
		{
			Item:       endpointWidget,
			Proportion: 1,
		},
		{
			Item:       buttons,
			Proportion: 1,
		},
	}
	main := widgets.Modal("Add a new endpoint", items)
	main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		switch pressedKey {
		case tcell.KeyEsc:
			page.onPopPage()
		case tcell.KeyEnter:
			page.addNewEndpoint()
		}
		return event
	})

	page.main = main
	page.collection = collectionsDropdown
	page.method = methodDropdown
	page.endpoint = endpointUrl

	return nil
}

func (page *NewEndpoint) addNewEndpoint() {
	// TODO: get collection id
	// collectionIndex, collection := page.collection.GetCurrentOption()

	_, selectedMethod := page.method.GetCurrentOption()
	endpoint := page.endpoint.GetText()

	err := page.onAddNewEndpoint(selectedMethod, endpoint, 0)
	if err != nil {
		page.onShowPopup(popup.POPUP_ERROR, fmt.Sprintf("Unable to add a new endpoint to collection due to %s", err))
		return
	}

	// TODO: update list of endpoints (UI)
}

func (page *NewEndpoint) Index() Index          { return NEW_ENDPOINT }
func (page *NewEndpoint) Page() tview.Primitive { return page.main }
