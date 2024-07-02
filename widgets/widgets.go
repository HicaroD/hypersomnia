package widgets

import (
	"github.com/HicaroD/hypersomnia/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InputField(placeholder string) *tview.InputField {
	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetFieldBackgroundColor(utils.COLOR_DARK_GREY)
	input.SetBackgroundColor(utils.COLOR_DARK_GREY)
	input.SetPlaceholder(placeholder)
	input.SetPlaceholderStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY))
	input.SetPlaceholderTextColor(tcell.ColorGrey)
	// TODO: set paste handler callback for validating input (for example, only
	// allow links to be pasted)
	return input
}

func TextArea(title string) *tview.TextArea {
	textArea := tview.NewTextArea()
	textArea.SetBorder(true)
	textArea.SetTitle(title)
	textArea.SetBackgroundColor(utils.COLOR_DARK_GREY)
	textArea.SetTextStyle(tcell.StyleDefault.Background(utils.COLOR_DARK_GREY))
	return textArea
}

func Dropdown(
	options []string,
	defaultOption int,
	selected func(text string, index int),
) *tview.DropDown {
	dropdown := tview.NewDropDown()
	dropdown.SetOptions(
		options,
		selected,
	)
	dropdown.SetCurrentOption(defaultOption)
	dropdown.SetFieldBackgroundColor(utils.COLOR_DARK_GREY)
	dropdown.SetBorder(true)
	dropdown.SetBackgroundColor(utils.COLOR_DARK_GREY)
	return dropdown
}

type PopupKind int

const (
	POPUP_ERROR PopupKind = iota
	POPUP_WARNING
)

// TODO: Don't make a new popup every time, make one and
// change the attributes every time you need a "new one"
func Popup(kind PopupKind, text string, onEsc func()) *tview.Flex {
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

	content := tview.NewTextView()
	content.SetTitle(title)
	content.SetText(text)
	content.SetBorder(true)
	content.SetBorderColor(borderColor)
	content.SetBackgroundColor(utils.COLOR_DARK_GREY)

	popup := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(content, 10, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)

	popup.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		switch pressedKey {
		case tcell.KeyESC:
			onEsc()
		}
		return event
	})

	return popup
}
