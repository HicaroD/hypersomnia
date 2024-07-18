package pages

import "github.com/rivo/tview"

type ListCollections struct {
	main *tview.Flex
}

func (page *ListCollections) Setup() error { return nil }

func (page *ListCollections) Index() Index          { return LIST_COLLECTIONS }
func (page *ListCollections) Page() tview.Primitive { return page.main }
