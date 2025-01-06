package pages

import (
	"fmt"

	"github.com/HicaroD/hypersomnia/models"
	"github.com/HicaroD/hypersomnia/utils"
	"github.com/HicaroD/hypersomnia/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListCollectionsPage struct {
	main                 *tview.Flex
	collectionListWidget *tview.List

	collections []*models.Collection

	onListCollections OnListCollectionsCallback
	onPopPage         OnPopPageCallback
}

func (page *ListCollectionsPage) Setup() error {
	collectionList := tview.NewList()
	collectionList.SetBackgroundColor(utils.COLOR_DARK_GREY)
	collectionList.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		selectedCollection := page.collections[index]
		fmt.Println(selectedCollection)
		// TODO: set the collection id in order to retrieve all endpoints
		// from this collection
	})
	page.collectionListWidget = collectionList
	err := page.UpdateCollectionList()
	if err != nil {
		return err
	}

	main := widgets.Modal("Collections", []widgets.Item{
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

func (page *ListCollectionsPage) UpdateCollectionList() error {
	collections, err := page.onListCollections()
	if err != nil {
		return err
	}
	page.collections = collections

	page.collectionListWidget = page.collectionListWidget.Clear()
	for _, collection := range page.collections {
		page.collectionListWidget.AddItem(collection.Name, "", 0, nil)
	}
	return nil
}

func (page *ListCollectionsPage) Index() Index          { return LIST_COLLECTIONS }
func (page *ListCollectionsPage) Page() tview.Primitive { return page.main }
