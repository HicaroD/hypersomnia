package widgets

import (
	"github.com/HicaroD/hypersomnia/models"
	"github.com/rivo/tview"
)

func HttpMethodsDropdown() *tview.DropDown {
	defaultOption := 0 // GET
	methodDropdown := Dropdown(models.ENDPOINT_METHODS, defaultOption, nil)
	return methodDropdown
}
