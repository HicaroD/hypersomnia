package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func HyperInputField(placeholder string) *tview.InputField {
	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetFieldBackgroundColor(DARK_GREY)
	input.SetBackgroundColor(DARK_GREY)
	input.SetPlaceholder(placeholder)
	input.SetPlaceholderStyle(tcell.StyleDefault.Background(DARK_GREY))
	input.SetPlaceholderTextColor(tcell.ColorGrey)
	// TODO: set paste handler callback for validating input (for example, only
	// allow links to be pasted)
	return input
}

func HyperTextArea(title string) *tview.TextArea {
	textArea := tview.NewTextArea()
	textArea.SetBorder(true)
	textArea.SetTitle(title)
	textArea.SetBackgroundColor(DARK_GREY)
	textArea.SetTextStyle(tcell.StyleDefault.Background(DARK_GREY))
	return textArea
}

func HyperDropdown(
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
	dropdown.SetFieldBackgroundColor(DARK_GREY)
	dropdown.SetBorder(true)
	dropdown.SetBackgroundColor(DARK_GREY)
	return dropdown
}

type PopupKind int

const (
	ERROR PopupKind = iota
	WARNING
)

func HyperPopup(kind PopupKind, text string) *tview.Flex {
	var title string
	var borderColor tcell.Color

	switch kind {
	case ERROR:
		title = "Error"
		borderColor = POPUP_RED
	case WARNING:
		title = "Warning"
		borderColor = POPUP_YELLOW
	}

	content := tview.NewTextView()
	content.SetTitle(title)
	content.SetText(text)
	content.SetBorder(true)
	content.SetBorderColor(borderColor)
	content.SetBackgroundColor(DARK_GREY)

	// escInfo := tview.NewTextView().SetTitle("Press ESC to close")

	popup := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(content, 10, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)

	return popup
}
