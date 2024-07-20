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

func TextButton(title string, backgroundColor tcell.Color, border bool, onPressed func()) *tview.Button {
	button := tview.NewButton(title)
	button.SetTitleColor(backgroundColor)
	button.SetStyle(tcell.StyleDefault.Background(backgroundColor))
	button.SetLabelColorActivated(utils.COLOR_BLACK)
	button.SetBorder(border)
	button.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		if action == tview.MouseLeftClick {
			onPressed()
		}
		return action, event
	})
	return button
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

type Item struct {
	Item       tview.Primitive
	Proportion int
}

func Modal(title string, items []Item) *tview.Flex {
	content := tview.NewFlex()
	content.SetDirection(tview.FlexRow)
	for _, item := range items {
		content.AddItem(item.Item, 0, item.Proportion, false)
	}
	content.SetBorder(true)
	content.SetBackgroundColor(utils.COLOR_DARK_GREY)

	modal := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(content, 0, 1, true).
			AddItem(nil, 0, 1, false), 0, 1, true).
		AddItem(nil, 0, 1, false)

	if title != "" {
		content.SetTitle(title)
	}

	return modal
}

func Text(t string) *tview.TextView {
	text := tview.NewTextView()
	text.SetText(t)
	text.SetBackgroundColor(utils.COLOR_DARK_GREY)
	return text
}

func layout(items []Item, row bool) *tview.Flex {
	content := tview.NewFlex()
	// NOTE: don't worry, it is not a mistake
	if row {
		content.SetDirection(tview.FlexColumn)
	} else {
		content.SetDirection(tview.FlexRow)
	}
	for _, item := range items {
		content.AddItem(item.Item, 0, item.Proportion, false)
	}
	return content
}

func Row(items []Item) *tview.Flex {
	return layout(items, true)
}

func Column(items []Item) *tview.Flex {
	return layout(items, false)
}
