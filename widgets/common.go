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
