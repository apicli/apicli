package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UrlInput struct {
	*tview.InputField
}

func NewUrlInput() *UrlInput {
	input := tview.NewInputField()
	input.SetLabel("URL: ")
	input.SetFieldBackgroundColor(tcell.ColorDefault)
	input.SetBackgroundColor(tcell.ColorDefault)
	return &UrlInput{
		InputField: input,
	}
}
