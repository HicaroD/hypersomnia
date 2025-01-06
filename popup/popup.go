package popup

import (
	"github.com/HicaroD/hypersomnia/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Kind int

const (
	POPUP_ERROR Kind = iota
	POPUP_WARNING
	POPUP_SUCCESS
)

type ShowPopupWidgetCallback func(tview.Primitive)

type Manager struct {
	onShowPopup ShowPopupWidgetCallback

	main    *tview.Flex
	content *tview.TextView
}

func New(onShowPopup ShowPopupWidgetCallback) *Manager {
	return &Manager{
		onShowPopup: onShowPopup,
	}
}

func (ppm *Manager) Setup() *tview.Flex {
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

func (ppm *Manager) Page() tview.Primitive { return ppm.main }

func (ppm *Manager) ShowPopup(kind Kind, text string) {
	var title string
	var borderColor tcell.Color

	switch kind {
	case POPUP_ERROR:
		title = "Error"
		borderColor = utils.COLOR_POPUP_RED
	case POPUP_WARNING:
		title = "Warning"
		borderColor = utils.COLOR_POPUP_YELLOW
	case POPUP_SUCCESS:
		title = "Success"
		borderColor = utils.COLOR_POPUP_GREEN
	}

	ppm.content.SetTitle(title)
	ppm.content.SetText(text)
	ppm.content.SetBorderColor(borderColor)

	ppm.onShowPopup(ppm.main)
}
