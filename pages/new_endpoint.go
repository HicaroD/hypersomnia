package pages

import (
	// "fmt"
	// "github.com/HicaroD/hypersomnia/popup"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewEndpoint struct {
	main *tview.Flex

	onPopPage OnPopPageCallback
}

func (page *NewEndpoint) Setup() error {
	// collectionsDropdown := widgets.Dropdown()

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
		// TODO: add dropdown for listing all collections
		// TODO: add dropdown for listing all possible HTTP methods
		// TODO: add text input field for the endpoint
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

	return nil
}

func (page *NewEndpoint) addNewEndpoint() {
	// TODO: add new endpoint to collection (database)
	// TODO: update list of endpoints (UI)
}

func (page *NewEndpoint) Index() Index          { return NEW_ENDPOINT }
func (page *NewEndpoint) Page() tview.Primitive { return page.main }
