package pages

import (
	"fmt"

	"github.com/HicaroD/hypersomnia/popup"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewCollection struct {
	main           *tview.Flex
	collectionName *tview.InputField

	onAddNewCollection OnAddNewCollectionCallback
	onPopPage          OnPopPageCallback
	onShowPopup        OnShowPopupCallback
}

func (page *NewCollection) Setup() error {
	collectionNameInput := widgets.InputField("My new collection")
	page.collectionName = collectionNameInput

	buttons := widgets.Row([]widgets.Item{
		{
			Item: widgets.TextButton(
				"Confirm",
				utils.COLOR_DARK_GREY,
				/* border= */ true,
				func() {
					name := page.collectionName.GetText()
					err := page.onAddNewCollection(name)
					if err != nil {
						page.onShowPopup(popup.POPUP_ERROR, fmt.Sprintf("Unable to create a new collection due to %s", err.Error()))
						return
					}
					page.onPopPage()
					page.onShowPopup(popup.POPUP_SUCCESS, "Collection was created successfully")
				},
			),
			Proportion: 1,
		},
		{
			Item: widgets.TextButton(
				"Cancel",
				utils.COLOR_DARK_GREY,
				/* border= */ true,
				func() {
					page.onPopPage()
				},
			),
			Proportion: 1,
		},
	})

	items := []widgets.Item{
		{
			Item:       widgets.Text("What is the name of your new collection?"),
			Proportion: 1,
		},
		{
			Item:       collectionNameInput,
			Proportion: 3,
		},
		{
			Item:       buttons,
			Proportion: 1,
		},
	}
	modal := widgets.Modal(items)
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		if pressedKey == tcell.KeyEsc {
			page.onPopPage()
		}
		return event
	})
	page.main = modal

	return nil
}

func (page *NewCollection) Index() Index          { return NEW_COLLECTION }
func (page *NewCollection) Page() tview.Primitive { return page.main }
