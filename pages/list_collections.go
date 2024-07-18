package pages

import (
	"fmt"

	"github.com/HicaroD/hypersomnia/utils"
	"github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListCollections struct {
	main *tview.Flex

	onListCollections OnListCollectionsCallback
	onPopPage         OnPopPageCallback
}

func (page *ListCollections) Setup() error {
	collections, err := page.onListCollections()
	if err != nil {
		return err
	}

	collectionList := tview.NewList()
	collectionList.SetBackgroundColor(utils.COLOR_DARK_GREY)

	for _, collection := range collections {
		collectionList.AddItem(collection.Name, "", 0, nil)
	}

	collectionList.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		selectedCollection := collections[index]
		fmt.Println(selectedCollection)
		// TODO: set the collection id in order to retrieve all endpoints
		// from this collection
	})

	main := widgets.Modal([]widgets.Item{
		{
			Item:       collectionList,
			Proportion: 1,
		},
	})

	main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		if pressedKey == tcell.KeyEsc {
			page.onPopPage()
		}
		return event
	})

	page.main = main
	return nil
}

func (page *ListCollections) Index() Index          { return LIST_COLLECTIONS }
func (page *ListCollections) Page() tview.Primitive { return page.main }
