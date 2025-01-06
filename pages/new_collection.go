package pages

import (
	"fmt"

	"github.com/HicaroD/hypersomnia/popup"
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewCollectionPage struct {
	main           *tview.Flex
	collectionName *tview.InputField

	onUpdateCollectionsList OnUpdateCollectionList
	onAddNewCollection      OnAddNewCollectionCallback
	onPopPage               OnPopPageCallback
	onShowPopup             OnShowPopupCallback
}

func (page *NewCollectionPage) Setup() error {
	collectionNameInput := widgets.InputField("My new collection")
	page.collectionName = collectionNameInput

	buttons := widgets.Row([]widgets.Item{
		{
			Item: widgets.TextButton(
				"Confirm",
				utils.COLOR_DARK_GREY,
				/* border= */ true,
				page.addNewCollection,
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
			Item:       collectionNameInput,
			Proportion: 4,
		},
		{
			Item:       buttons,
			Proportion: 1,
		},
	}
	main := widgets.Modal("Add a new collection", items)
	main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		switch pressedKey {
		case tcell.KeyEsc:
			page.onPopPage()
		case tcell.KeyEnter:
			page.addNewCollection()
		}
		return event
	})
	page.main = main

	return nil
}

func (page *NewCollectionPage) addNewCollection() {
	var err error

	name := page.collectionName.GetText()
	err = page.onAddNewCollection(name)
	if err != nil {
		page.onShowPopup(popup.POPUP_ERROR, fmt.Sprintf("Unable to create a new collection due to %s", err.Error()))
		return
	}

	page.onPopPage()
	page.onShowPopup(popup.POPUP_SUCCESS, "Collection was created successfully")

	err = page.onUpdateCollectionsList()
	if err != nil {
		page.onShowPopup(popup.POPUP_ERROR, fmt.Sprintf("Unable to update collection list due to %s", err.Error()))
		return
	}
}

func (page *NewCollectionPage) Index() Index          { return NEW_COLLECTION }
func (page *NewCollectionPage) Page() tview.Primitive { return page.main }
