package pages

import (
	utils "github.com/HicaroD/hypersomnia/utils"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/rivo/tview"
)

type (
	OnAddNewCollectionCallback func(name string) error
	OnPopPageCallback          func()
)

type NewCollection struct {
	main           *tview.Flex
	collectionName *tview.InputField

	onAddNewCollection OnAddNewCollectionCallback
	onPopPage          OnPopPageCallback
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
						// TODO: show popup here as well
						panic(err)
					}
					page.onPopPage()
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
	page.main = modal

	return nil
}

func (page *NewCollection) Index() Index          { return NEW_COLLECTION }
func (page *NewCollection) Page() tview.Primitive { return page.main }
