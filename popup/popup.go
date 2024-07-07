package popup

import (
	"github.com/HicaroD/hypersomnia/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PopupKind int

const (
	POPUP_ERROR PopupKind = iota
	POPUP_WARNING
)

type ShowPopupWidgetCallback func(tview.Primitive)

type PopupManager struct {
	OnShowPopup ShowPopupWidgetCallback

	main    *tview.Flex
	content *tview.TextView
}

func (ppm *PopupManager) Setup() *tview.Flex {
	content := tview.NewTextView()
	content.SetBorder(true)
	content.SetBackgroundColor(utils.COLOR_DARK_GREY)

	popup := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(content, 10, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)

	ppm.content = content
	ppm.main = popup

	return popup
}

func (ppm *PopupManager) Page() tview.Primitive { return ppm.main }

func (ppm *PopupManager) ShowPopup(kind PopupKind, text string) {
	var title string
	var borderColor tcell.Color

	switch kind {
	case POPUP_ERROR:
		title = "Error"
		borderColor = utils.COLOR_POPUP_RED
	case POPUP_WARNING:
		title = "Warning"
		borderColor = utils.COLOR_POPUP_YELLOW
	}

	ppm.content.SetTitle(title)
	ppm.content.SetText(text)
	ppm.content.SetBorderColor(borderColor)

	ppm.OnShowPopup(ppm.main)
}
